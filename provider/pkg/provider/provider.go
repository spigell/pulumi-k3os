// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jpillora/backoff"
	"github.com/pkg/errors"

	"github.com/pulumi/pulumi/pkg/v3/resource/provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
	rpc "github.com/pulumi/pulumi/sdk/v3/proto/go"

	pbempty "github.com/golang/protobuf/ptypes/empty"

	"github.com/spigell/pulumi-k3os/provider/pkg/resources"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var waitForCompletionMsg = "Waiting for the reboot and reconnection"

type k3osProvider struct {
	host           *provider.HostClient
	name           string
	version        string
	yamlRenderMode bool
	yamlDirectory  string
}

func makeProvider(host *provider.HostClient, name, version string) (rpc.ResourceProviderServer, error) {
	// Return the new provider
	return &k3osProvider{
		host:    host,
		name:    name,
		version: version,
	}, nil
}

// CheckConfig validates the configuration for this provider.
func (k *k3osProvider) CheckConfig(ctx context.Context, req *rpc.CheckRequest) (*rpc.CheckResponse, error) {
	return &rpc.CheckResponse{Inputs: req.GetNews()}, nil
}

// DiffConfig diffs the configuration for this provider.
func (k *k3osProvider) DiffConfig(ctx context.Context, req *rpc.DiffRequest) (*rpc.DiffResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.DiffConfig(%s)", k.name, urn)
	logging.V(9).Infof("%s executing", label)

	olds, err := plugin.UnmarshalProperties(req.GetOlds(), plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.olds", label),
		KeepUnknowns: true,
		SkipNulls:    true,
	})
	if err != nil {
		return nil, err
	}
	news, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.news", label),
		KeepUnknowns: true,
		SkipNulls:    true,
		RejectAssets: true,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "DiffConfig failed because of malformed resource inputs")
	}

	var diffs, replaces []string

	if olds["renderYamlToDirectory"] != news["renderYamlToDirectory"] {
		// If the render directory changes, the config will be replaced.
		diffs = append(diffs, "renderYamlToDirectory")
		replaces = append(replaces, "renderYamlToDirectory")
	}

	if len(replaces) > 0 {
		return &rpc.DiffResponse{
			Changes:  rpc.DiffResponse_DIFF_SOME,
			Diffs:    diffs,
			Replaces: replaces,
		}, nil
	}

	return &rpc.DiffResponse{
		Changes: rpc.DiffResponse_DIFF_NONE,
	}, nil
}

// Configure configures the resource provider with "globals" that control its behavior.
func (k *k3osProvider) Configure(_ context.Context, req *rpc.ConfigureRequest) (*rpc.ConfigureResponse, error) {
	vars := req.GetVariables()

	renderYamlToDirectory := func() string {
		if directory, exists := vars["k3os:config:renderYamlToDirectory"]; exists && directory != "" {
			return directory
		}
		return ""
	}
	k.yamlDirectory = renderYamlToDirectory()
	k.yamlRenderMode = len(k.yamlDirectory) > 0

	return &rpc.ConfigureResponse{
		AcceptSecrets:   true,
		SupportsPreview: false,
	}, nil
}

// Invoke dynamically executes a built-in function in the provider.
func (k *k3osProvider) Invoke(_ context.Context, req *rpc.InvokeRequest) (*rpc.InvokeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Invoke is not yet implemented")
}

// StreamInvoke dynamically executes a built-in function in the provider. The result is streamed
// back as a series of messages.
func (k *k3osProvider) StreamInvoke(req *rpc.InvokeRequest, server rpc.ResourceProvider_StreamInvokeServer) error {
	return status.Error(codes.Unimplemented, "StreamInvoke is not yet implemented")
}

// Check validates that the given property bag is valid for a resource of the given type and returns
// the inputs that should be passed to successive calls to Diff, Create, or Update for this
// resource. As a rule, the provider inputs returned by a call to Check should preserve the original
// representation of the properties as present in the program inputs. Though this rule is not
// required for correctness, violations thereof can negatively impact the end-user experience, as
// the provider inputs are using for detecting and rendering diffs.
func (k *k3osProvider) Check(ctx context.Context, req *rpc.CheckRequest) (*rpc.CheckResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Check(%s)", k.name, urn)
	logging.V(9).Infof("%s executing", label)

	newResInputs := req.GetNews()

	return &rpc.CheckResponse{Inputs: newResInputs}, nil
}

// Diff checks what impacts a hypothetical update will have on the resource's properties.
func (k *k3osProvider) Diff(ctx context.Context, req *rpc.DiffRequest) (*rpc.DiffResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Diff(%s)", k.name, urn)
	logging.V(9).Infof("%s executing", label)

	oldState, err := plugin.UnmarshalProperties(req.GetOlds(), plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.oldState", label),
		KeepUnknowns: true,
		KeepSecrets:  true,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "diff failed because malformed resource inputs")
	}

	oldInputs := parseCheckpointObject(oldState)

	newInputs, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.newInputs", label),
		KeepUnknowns: true,
		KeepSecrets:  true,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "diff failed because malformed resource inputs")
	}

	var replaces []string

	diff := oldInputs.Diff(newInputs)
	if diff == nil {
		return &rpc.DiffResponse{Changes: rpc.DiffResponse_DIFF_NONE}, nil
	}

	return &rpc.DiffResponse{Changes: rpc.DiffResponse_DIFF_SOME, Replaces: replaces}, nil
}

// Create allocates a new instance of the provided resource and returns its unique ID afterwards.
func (k *k3osProvider) Create(ctx context.Context, req *rpc.CreateRequest) (*rpc.CreateResponse, error) {
	inputs, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	urn := resource.URN(req.GetUrn())

	if k.yamlRenderMode {
		err := renderYaml(k.yamlDirectory, inputs["nodeConfiguration"])
		if err != nil {
			return nil, err
		}

		outputProperties, err := plugin.MarshalProperties(
			checkpointObject(inputs, make(map[string]interface{})),
			plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true},
		)
		if err != nil {
			return nil, err
		}

		return &rpc.CreateResponse{
			Id:         string(resource.URN(req.GetUrn())),
			Properties: outputProperties,
		}, nil
	}

	_, err = resources.SetNodeConfig(inputs)
	if err != nil {
		return nil, err
	}

	_ = k.host.LogStatus(ctx, diag.Info, urn, fmt.Sprintf("Create: %s", waitForCompletionMsg))
	resp := k.waitForResourceOpCompletion(inputs["connection"])
	_ = k.host.LogStatus(ctx, diag.Info, urn, "Create: successful")

	outputProperties, err := plugin.MarshalProperties(
		checkpointObject(inputs, resp),
		plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true},
	)
	if err != nil {
		return nil, err
	}

	return &rpc.CreateResponse{
		Id:         string(urn),
		Properties: outputProperties,
	}, nil
}

func (k *k3osProvider) waitForResourceOpCompletion(conn resource.PropertyValue) map[string]interface{} {
	retryPolicy := backoff.Backoff{
		Min:    1 * time.Second,
		Max:    15 * time.Second,
		Factor: 1.5,
		Jitter: true,
	}
	time.Sleep(2 * time.Second)

	for {
		resp, _ := resources.CheckNodeConnection(conn)

		status, hasStatus := resp["success"].(bool)
		if completed := (hasStatus && status); completed {
			return resp
		}

		time.Sleep(retryPolicy.Duration())
	}
}

// Read the current live state associated with a resource.
func (k *k3osProvider) Read(ctx context.Context, req *rpc.ReadRequest) (*rpc.ReadResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Read/Refresh/Import is not yet implemented")
}

// Update updates an existing resource with new values.
func (k *k3osProvider) Update(ctx context.Context, req *rpc.UpdateRequest) (*rpc.UpdateResponse, error) {
	urn := resource.URN(req.GetUrn())
	news, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	if k.yamlRenderMode {
		err := renderYaml(k.yamlDirectory, news["nodeConfiguration"])
		if err != nil {
			return nil, err
		}

		outputProperties, err := plugin.MarshalProperties(
			checkpointObject(news, make(map[string]interface{})),
			plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true},
		)
		if err != nil {
			return nil, err
		}

		return &rpc.UpdateResponse{
			Properties: outputProperties,
		}, nil
	}

	_, err = resources.SetNodeConfig(news)
	if err != nil {
		return nil, err
	}

	_ = k.host.LogStatus(ctx, diag.Info, urn, fmt.Sprintf("Update: %s", waitForCompletionMsg))
	resp := k.waitForResourceOpCompletion(news["connection"])
	_ = k.host.LogStatus(ctx, diag.Info, urn, "Update: successful")

	outputProperties, err := plugin.MarshalProperties(
		checkpointObject(news, resp),
		plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true},
	)
	if err != nil {
		return nil, err
	}

	return &rpc.UpdateResponse{
		Properties: outputProperties,
	}, nil
}

// Delete tears down an existing resource with the given ID.  If it fails, the resource is assumed
// to still exist.
func (k *k3osProvider) Delete(ctx context.Context, req *rpc.DeleteRequest) (*pbempty.Empty, error) {
	urn := resource.URN(req.GetUrn())
	props, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	inputs := parseCheckpointObject(props)

	if k.yamlRenderMode {
		err := os.RemoveAll(k.yamlDirectory)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to remove YAML dir: %q", k.yamlDirectory)
		}
	} else {
		err = resources.DeleteNodeConfig(inputs["connection"])
		if err != nil {
			return nil, err
		}

		_ = k.host.LogStatus(ctx, diag.Info, urn, fmt.Sprintf("Delete: %s", waitForCompletionMsg))
		_ = k.waitForResourceOpCompletion(inputs["connection"])
	}

	return &pbempty.Empty{}, nil
}

// Construct creates a new component resource.
func (k *k3osProvider) Construct(_ context.Context, _ *rpc.ConstructRequest) (*rpc.ConstructResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Construct is not yet implemented")
}

// GetPluginInfo returns generic information about this plugin, like its version.
func (k *k3osProvider) GetPluginInfo(context.Context, *pbempty.Empty) (*rpc.PluginInfo, error) {
	return &rpc.PluginInfo{
		Version: k.version,
	}, nil
}

// GetSchema returns the JSON-serialized schema for the provider.
func (k *k3osProvider) GetSchema(_ context.Context, req *rpc.GetSchemaRequest) (*rpc.GetSchemaResponse, error) {
	return &rpc.GetSchemaResponse{}, nil
}

// Cancel signals the provider to gracefully shut down and abort any ongoing resource operations.
// Operations aborted in this way will return an error (e.g., `Update` and `Create` will either a
// creation error or an initialization error). Since Cancel is advisory and non-blocking, it is up
// to the host to decide how long to wait after Cancel is called before (e.g.)
// hard-closing any gRPC connection.
func (k *k3osProvider) Cancel(context.Context, *pbempty.Empty) (*pbempty.Empty, error) {
	return &pbempty.Empty{}, nil
}

func checkpointObject(inputs resource.PropertyMap, outputs map[string]interface{}) resource.PropertyMap {
	object := resource.NewPropertyMapFromMap(outputs)
	object["__inputs"] = resource.MakeSecret(resource.NewObjectProperty(inputs))
	return object
}

func parseCheckpointObject(obj resource.PropertyMap) resource.PropertyMap {
	if inputs, ok := obj["__inputs"]; ok {
		return inputs.ObjectValue()
	}

	return nil
}

func renderYaml(dir string, r resource.PropertyValue) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0700)
		if err != nil {
			return errors.Wrapf(err, "failed to create directory for rendered YAML: %q", dir)
		}

		path, err := resources.PutNodeConfigToDir(dir, r)
		if err != nil {
			return errors.Wrapf(err, "failed to render YAML to: %q", path)
		}
	}

	return nil
}
