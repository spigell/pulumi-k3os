// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package k3os

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CloudInitFiles struct {
	Content            *string `pulumi:"content"`
	Encoding           *string `pulumi:"encoding"`
	Owner              *string `pulumi:"owner"`
	Path               *string `pulumi:"path"`
	RawFilePermissions *string `pulumi:"rawFilePermissions"`
}

// CloudInitFilesInput is an input type that accepts CloudInitFilesArgs and CloudInitFilesOutput values.
// You can construct a concrete instance of `CloudInitFilesInput` via:
//
//          CloudInitFilesArgs{...}
type CloudInitFilesInput interface {
	pulumi.Input

	ToCloudInitFilesOutput() CloudInitFilesOutput
	ToCloudInitFilesOutputWithContext(context.Context) CloudInitFilesOutput
}

type CloudInitFilesArgs struct {
	Content            pulumi.StringPtrInput `pulumi:"content"`
	Encoding           pulumi.StringPtrInput `pulumi:"encoding"`
	Owner              pulumi.StringPtrInput `pulumi:"owner"`
	Path               pulumi.StringPtrInput `pulumi:"path"`
	RawFilePermissions pulumi.StringPtrInput `pulumi:"rawFilePermissions"`
}

func (CloudInitFilesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudInitFiles)(nil)).Elem()
}

func (i CloudInitFilesArgs) ToCloudInitFilesOutput() CloudInitFilesOutput {
	return i.ToCloudInitFilesOutputWithContext(context.Background())
}

func (i CloudInitFilesArgs) ToCloudInitFilesOutputWithContext(ctx context.Context) CloudInitFilesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudInitFilesOutput)
}

func (i CloudInitFilesArgs) ToCloudInitFilesPtrOutput() CloudInitFilesPtrOutput {
	return i.ToCloudInitFilesPtrOutputWithContext(context.Background())
}

func (i CloudInitFilesArgs) ToCloudInitFilesPtrOutputWithContext(ctx context.Context) CloudInitFilesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudInitFilesOutput).ToCloudInitFilesPtrOutputWithContext(ctx)
}

// CloudInitFilesPtrInput is an input type that accepts CloudInitFilesArgs, CloudInitFilesPtr and CloudInitFilesPtrOutput values.
// You can construct a concrete instance of `CloudInitFilesPtrInput` via:
//
//          CloudInitFilesArgs{...}
//
//  or:
//
//          nil
type CloudInitFilesPtrInput interface {
	pulumi.Input

	ToCloudInitFilesPtrOutput() CloudInitFilesPtrOutput
	ToCloudInitFilesPtrOutputWithContext(context.Context) CloudInitFilesPtrOutput
}

type cloudInitFilesPtrType CloudInitFilesArgs

func CloudInitFilesPtr(v *CloudInitFilesArgs) CloudInitFilesPtrInput {
	return (*cloudInitFilesPtrType)(v)
}

func (*cloudInitFilesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudInitFiles)(nil)).Elem()
}

func (i *cloudInitFilesPtrType) ToCloudInitFilesPtrOutput() CloudInitFilesPtrOutput {
	return i.ToCloudInitFilesPtrOutputWithContext(context.Background())
}

func (i *cloudInitFilesPtrType) ToCloudInitFilesPtrOutputWithContext(ctx context.Context) CloudInitFilesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudInitFilesPtrOutput)
}

type CloudInitFilesOutput struct{ *pulumi.OutputState }

func (CloudInitFilesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudInitFiles)(nil)).Elem()
}

func (o CloudInitFilesOutput) ToCloudInitFilesOutput() CloudInitFilesOutput {
	return o
}

func (o CloudInitFilesOutput) ToCloudInitFilesOutputWithContext(ctx context.Context) CloudInitFilesOutput {
	return o
}

func (o CloudInitFilesOutput) ToCloudInitFilesPtrOutput() CloudInitFilesPtrOutput {
	return o.ToCloudInitFilesPtrOutputWithContext(context.Background())
}

func (o CloudInitFilesOutput) ToCloudInitFilesPtrOutputWithContext(ctx context.Context) CloudInitFilesPtrOutput {
	return o.ApplyT(func(v CloudInitFiles) *CloudInitFiles {
		return &v
	}).(CloudInitFilesPtrOutput)
}
func (o CloudInitFilesOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudInitFiles) *string { return v.Content }).(pulumi.StringPtrOutput)
}

func (o CloudInitFilesOutput) Encoding() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudInitFiles) *string { return v.Encoding }).(pulumi.StringPtrOutput)
}

func (o CloudInitFilesOutput) Owner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudInitFiles) *string { return v.Owner }).(pulumi.StringPtrOutput)
}

func (o CloudInitFilesOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudInitFiles) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o CloudInitFilesOutput) RawFilePermissions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudInitFiles) *string { return v.RawFilePermissions }).(pulumi.StringPtrOutput)
}

type CloudInitFilesPtrOutput struct{ *pulumi.OutputState }

func (CloudInitFilesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudInitFiles)(nil)).Elem()
}

func (o CloudInitFilesPtrOutput) ToCloudInitFilesPtrOutput() CloudInitFilesPtrOutput {
	return o
}

func (o CloudInitFilesPtrOutput) ToCloudInitFilesPtrOutputWithContext(ctx context.Context) CloudInitFilesPtrOutput {
	return o
}

func (o CloudInitFilesPtrOutput) Elem() CloudInitFilesOutput {
	return o.ApplyT(func(v *CloudInitFiles) CloudInitFiles { return *v }).(CloudInitFilesOutput)
}

func (o CloudInitFilesPtrOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudInitFiles) *string {
		if v == nil {
			return nil
		}
		return v.Content
	}).(pulumi.StringPtrOutput)
}

func (o CloudInitFilesPtrOutput) Encoding() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudInitFiles) *string {
		if v == nil {
			return nil
		}
		return v.Encoding
	}).(pulumi.StringPtrOutput)
}

func (o CloudInitFilesPtrOutput) Owner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudInitFiles) *string {
		if v == nil {
			return nil
		}
		return v.Owner
	}).(pulumi.StringPtrOutput)
}

func (o CloudInitFilesPtrOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudInitFiles) *string {
		if v == nil {
			return nil
		}
		return v.Path
	}).(pulumi.StringPtrOutput)
}

func (o CloudInitFilesPtrOutput) RawFilePermissions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudInitFiles) *string {
		if v == nil {
			return nil
		}
		return v.RawFilePermissions
	}).(pulumi.StringPtrOutput)
}

type Connection struct {
	Addr *string `pulumi:"addr"`
	Key  *string `pulumi:"key"`
	User *string `pulumi:"user"`
}

// ConnectionInput is an input type that accepts ConnectionArgs and ConnectionOutput values.
// You can construct a concrete instance of `ConnectionInput` via:
//
//          ConnectionArgs{...}
type ConnectionInput interface {
	pulumi.Input

	ToConnectionOutput() ConnectionOutput
	ToConnectionOutputWithContext(context.Context) ConnectionOutput
}

type ConnectionArgs struct {
	Addr pulumi.StringPtrInput `pulumi:"addr"`
	Key  pulumi.StringPtrInput `pulumi:"key"`
	User pulumi.StringPtrInput `pulumi:"user"`
}

func (ConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Connection)(nil)).Elem()
}

func (i ConnectionArgs) ToConnectionOutput() ConnectionOutput {
	return i.ToConnectionOutputWithContext(context.Background())
}

func (i ConnectionArgs) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionOutput)
}

func (i ConnectionArgs) ToConnectionPtrOutput() ConnectionPtrOutput {
	return i.ToConnectionPtrOutputWithContext(context.Background())
}

func (i ConnectionArgs) ToConnectionPtrOutputWithContext(ctx context.Context) ConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionOutput).ToConnectionPtrOutputWithContext(ctx)
}

// ConnectionPtrInput is an input type that accepts ConnectionArgs, ConnectionPtr and ConnectionPtrOutput values.
// You can construct a concrete instance of `ConnectionPtrInput` via:
//
//          ConnectionArgs{...}
//
//  or:
//
//          nil
type ConnectionPtrInput interface {
	pulumi.Input

	ToConnectionPtrOutput() ConnectionPtrOutput
	ToConnectionPtrOutputWithContext(context.Context) ConnectionPtrOutput
}

type connectionPtrType ConnectionArgs

func ConnectionPtr(v *ConnectionArgs) ConnectionPtrInput {
	return (*connectionPtrType)(v)
}

func (*connectionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Connection)(nil)).Elem()
}

func (i *connectionPtrType) ToConnectionPtrOutput() ConnectionPtrOutput {
	return i.ToConnectionPtrOutputWithContext(context.Background())
}

func (i *connectionPtrType) ToConnectionPtrOutputWithContext(ctx context.Context) ConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionPtrOutput)
}

type ConnectionOutput struct{ *pulumi.OutputState }

func (ConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Connection)(nil)).Elem()
}

func (o ConnectionOutput) ToConnectionOutput() ConnectionOutput {
	return o
}

func (o ConnectionOutput) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return o
}

func (o ConnectionOutput) ToConnectionPtrOutput() ConnectionPtrOutput {
	return o.ToConnectionPtrOutputWithContext(context.Background())
}

func (o ConnectionOutput) ToConnectionPtrOutputWithContext(ctx context.Context) ConnectionPtrOutput {
	return o.ApplyT(func(v Connection) *Connection {
		return &v
	}).(ConnectionPtrOutput)
}
func (o ConnectionOutput) Addr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Connection) *string { return v.Addr }).(pulumi.StringPtrOutput)
}

func (o ConnectionOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Connection) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o ConnectionOutput) User() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Connection) *string { return v.User }).(pulumi.StringPtrOutput)
}

type ConnectionPtrOutput struct{ *pulumi.OutputState }

func (ConnectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Connection)(nil)).Elem()
}

func (o ConnectionPtrOutput) ToConnectionPtrOutput() ConnectionPtrOutput {
	return o
}

func (o ConnectionPtrOutput) ToConnectionPtrOutputWithContext(ctx context.Context) ConnectionPtrOutput {
	return o
}

func (o ConnectionPtrOutput) Elem() ConnectionOutput {
	return o.ApplyT(func(v *Connection) Connection { return *v }).(ConnectionOutput)
}

func (o ConnectionPtrOutput) Addr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Connection) *string {
		if v == nil {
			return nil
		}
		return v.Addr
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionPtrOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Connection) *string {
		if v == nil {
			return nil
		}
		return v.Key
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionPtrOutput) User() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Connection) *string {
		if v == nil {
			return nil
		}
		return v.User
	}).(pulumi.StringPtrOutput)
}

type K3OS struct {
	Environment map[string]string `pulumi:"environment"`
	K3sArgs     []string          `pulumi:"k3sArgs"`
	Labels      map[string]string `pulumi:"labels"`
	Password    *string           `pulumi:"password"`
	ServerUrl   *string           `pulumi:"serverUrl"`
	Taints      []string          `pulumi:"taints"`
	Token       *string           `pulumi:"token"`
}

// K3OSInput is an input type that accepts K3OSArgs and K3OSOutput values.
// You can construct a concrete instance of `K3OSInput` via:
//
//          K3OSArgs{...}
type K3OSInput interface {
	pulumi.Input

	ToK3OSOutput() K3OSOutput
	ToK3OSOutputWithContext(context.Context) K3OSOutput
}

type K3OSArgs struct {
	Environment pulumi.StringMapInput   `pulumi:"environment"`
	K3sArgs     pulumi.StringArrayInput `pulumi:"k3sArgs"`
	Labels      pulumi.StringMapInput   `pulumi:"labels"`
	Password    pulumi.StringPtrInput   `pulumi:"password"`
	ServerUrl   pulumi.StringPtrInput   `pulumi:"serverUrl"`
	Taints      pulumi.StringArrayInput `pulumi:"taints"`
	Token       pulumi.StringPtrInput   `pulumi:"token"`
}

func (K3OSArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*K3OS)(nil)).Elem()
}

func (i K3OSArgs) ToK3OSOutput() K3OSOutput {
	return i.ToK3OSOutputWithContext(context.Background())
}

func (i K3OSArgs) ToK3OSOutputWithContext(ctx context.Context) K3OSOutput {
	return pulumi.ToOutputWithContext(ctx, i).(K3OSOutput)
}

func (i K3OSArgs) ToK3OSPtrOutput() K3OSPtrOutput {
	return i.ToK3OSPtrOutputWithContext(context.Background())
}

func (i K3OSArgs) ToK3OSPtrOutputWithContext(ctx context.Context) K3OSPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(K3OSOutput).ToK3OSPtrOutputWithContext(ctx)
}

// K3OSPtrInput is an input type that accepts K3OSArgs, K3OSPtr and K3OSPtrOutput values.
// You can construct a concrete instance of `K3OSPtrInput` via:
//
//          K3OSArgs{...}
//
//  or:
//
//          nil
type K3OSPtrInput interface {
	pulumi.Input

	ToK3OSPtrOutput() K3OSPtrOutput
	ToK3OSPtrOutputWithContext(context.Context) K3OSPtrOutput
}

type k3osPtrType K3OSArgs

func K3OSPtr(v *K3OSArgs) K3OSPtrInput {
	return (*k3osPtrType)(v)
}

func (*k3osPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**K3OS)(nil)).Elem()
}

func (i *k3osPtrType) ToK3OSPtrOutput() K3OSPtrOutput {
	return i.ToK3OSPtrOutputWithContext(context.Background())
}

func (i *k3osPtrType) ToK3OSPtrOutputWithContext(ctx context.Context) K3OSPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(K3OSPtrOutput)
}

type K3OSOutput struct{ *pulumi.OutputState }

func (K3OSOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*K3OS)(nil)).Elem()
}

func (o K3OSOutput) ToK3OSOutput() K3OSOutput {
	return o
}

func (o K3OSOutput) ToK3OSOutputWithContext(ctx context.Context) K3OSOutput {
	return o
}

func (o K3OSOutput) ToK3OSPtrOutput() K3OSPtrOutput {
	return o.ToK3OSPtrOutputWithContext(context.Background())
}

func (o K3OSOutput) ToK3OSPtrOutputWithContext(ctx context.Context) K3OSPtrOutput {
	return o.ApplyT(func(v K3OS) *K3OS {
		return &v
	}).(K3OSPtrOutput)
}
func (o K3OSOutput) Environment() pulumi.StringMapOutput {
	return o.ApplyT(func(v K3OS) map[string]string { return v.Environment }).(pulumi.StringMapOutput)
}

func (o K3OSOutput) K3sArgs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v K3OS) []string { return v.K3sArgs }).(pulumi.StringArrayOutput)
}

func (o K3OSOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v K3OS) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

func (o K3OSOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v K3OS) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o K3OSOutput) ServerUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v K3OS) *string { return v.ServerUrl }).(pulumi.StringPtrOutput)
}

func (o K3OSOutput) Taints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v K3OS) []string { return v.Taints }).(pulumi.StringArrayOutput)
}

func (o K3OSOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v K3OS) *string { return v.Token }).(pulumi.StringPtrOutput)
}

type K3OSPtrOutput struct{ *pulumi.OutputState }

func (K3OSPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**K3OS)(nil)).Elem()
}

func (o K3OSPtrOutput) ToK3OSPtrOutput() K3OSPtrOutput {
	return o
}

func (o K3OSPtrOutput) ToK3OSPtrOutputWithContext(ctx context.Context) K3OSPtrOutput {
	return o
}

func (o K3OSPtrOutput) Elem() K3OSOutput {
	return o.ApplyT(func(v *K3OS) K3OS { return *v }).(K3OSOutput)
}

func (o K3OSPtrOutput) Environment() pulumi.StringMapOutput {
	return o.ApplyT(func(v *K3OS) map[string]string {
		if v == nil {
			return nil
		}
		return v.Environment
	}).(pulumi.StringMapOutput)
}

func (o K3OSPtrOutput) K3sArgs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *K3OS) []string {
		if v == nil {
			return nil
		}
		return v.K3sArgs
	}).(pulumi.StringArrayOutput)
}

func (o K3OSPtrOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *K3OS) map[string]string {
		if v == nil {
			return nil
		}
		return v.Labels
	}).(pulumi.StringMapOutput)
}

func (o K3OSPtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *K3OS) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o K3OSPtrOutput) ServerUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *K3OS) *string {
		if v == nil {
			return nil
		}
		return v.ServerUrl
	}).(pulumi.StringPtrOutput)
}

func (o K3OSPtrOutput) Taints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *K3OS) []string {
		if v == nil {
			return nil
		}
		return v.Taints
	}).(pulumi.StringArrayOutput)
}

func (o K3OSPtrOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *K3OS) *string {
		if v == nil {
			return nil
		}
		return v.Token
	}).(pulumi.StringPtrOutput)
}

type NodeConfiguration struct {
	BootCmd           []string          `pulumi:"bootCmd"`
	Datasources       []string          `pulumi:"datasources"`
	Hostname          *string           `pulumi:"hostname"`
	InitCmd           []string          `pulumi:"initCmd"`
	K3OS              *K3OS             `pulumi:"k3OS"`
	Modules           []string          `pulumi:"modules"`
	RunCmd            []string          `pulumi:"runCmd"`
	SshAuthorizerKeys []string          `pulumi:"sshAuthorizerKeys"`
	Sysctls           map[string]string `pulumi:"sysctls"`
	WriteFiles        *CloudInitFiles   `pulumi:"writeFiles"`
}

// NodeConfigurationInput is an input type that accepts NodeConfigurationArgs and NodeConfigurationOutput values.
// You can construct a concrete instance of `NodeConfigurationInput` via:
//
//          NodeConfigurationArgs{...}
type NodeConfigurationInput interface {
	pulumi.Input

	ToNodeConfigurationOutput() NodeConfigurationOutput
	ToNodeConfigurationOutputWithContext(context.Context) NodeConfigurationOutput
}

type NodeConfigurationArgs struct {
	BootCmd           pulumi.StringArrayInput `pulumi:"bootCmd"`
	Datasources       pulumi.StringArrayInput `pulumi:"datasources"`
	Hostname          pulumi.StringPtrInput   `pulumi:"hostname"`
	InitCmd           pulumi.StringArrayInput `pulumi:"initCmd"`
	K3OS              K3OSPtrInput            `pulumi:"k3OS"`
	Modules           pulumi.StringArrayInput `pulumi:"modules"`
	RunCmd            pulumi.StringArrayInput `pulumi:"runCmd"`
	SshAuthorizerKeys pulumi.StringArrayInput `pulumi:"sshAuthorizerKeys"`
	Sysctls           pulumi.StringMapInput   `pulumi:"sysctls"`
	WriteFiles        CloudInitFilesPtrInput  `pulumi:"writeFiles"`
}

func (NodeConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeConfiguration)(nil)).Elem()
}

func (i NodeConfigurationArgs) ToNodeConfigurationOutput() NodeConfigurationOutput {
	return i.ToNodeConfigurationOutputWithContext(context.Background())
}

func (i NodeConfigurationArgs) ToNodeConfigurationOutputWithContext(ctx context.Context) NodeConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeConfigurationOutput)
}

func (i NodeConfigurationArgs) ToNodeConfigurationPtrOutput() NodeConfigurationPtrOutput {
	return i.ToNodeConfigurationPtrOutputWithContext(context.Background())
}

func (i NodeConfigurationArgs) ToNodeConfigurationPtrOutputWithContext(ctx context.Context) NodeConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeConfigurationOutput).ToNodeConfigurationPtrOutputWithContext(ctx)
}

// NodeConfigurationPtrInput is an input type that accepts NodeConfigurationArgs, NodeConfigurationPtr and NodeConfigurationPtrOutput values.
// You can construct a concrete instance of `NodeConfigurationPtrInput` via:
//
//          NodeConfigurationArgs{...}
//
//  or:
//
//          nil
type NodeConfigurationPtrInput interface {
	pulumi.Input

	ToNodeConfigurationPtrOutput() NodeConfigurationPtrOutput
	ToNodeConfigurationPtrOutputWithContext(context.Context) NodeConfigurationPtrOutput
}

type nodeConfigurationPtrType NodeConfigurationArgs

func NodeConfigurationPtr(v *NodeConfigurationArgs) NodeConfigurationPtrInput {
	return (*nodeConfigurationPtrType)(v)
}

func (*nodeConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NodeConfiguration)(nil)).Elem()
}

func (i *nodeConfigurationPtrType) ToNodeConfigurationPtrOutput() NodeConfigurationPtrOutput {
	return i.ToNodeConfigurationPtrOutputWithContext(context.Background())
}

func (i *nodeConfigurationPtrType) ToNodeConfigurationPtrOutputWithContext(ctx context.Context) NodeConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeConfigurationPtrOutput)
}

type NodeConfigurationOutput struct{ *pulumi.OutputState }

func (NodeConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeConfiguration)(nil)).Elem()
}

func (o NodeConfigurationOutput) ToNodeConfigurationOutput() NodeConfigurationOutput {
	return o
}

func (o NodeConfigurationOutput) ToNodeConfigurationOutputWithContext(ctx context.Context) NodeConfigurationOutput {
	return o
}

func (o NodeConfigurationOutput) ToNodeConfigurationPtrOutput() NodeConfigurationPtrOutput {
	return o.ToNodeConfigurationPtrOutputWithContext(context.Background())
}

func (o NodeConfigurationOutput) ToNodeConfigurationPtrOutputWithContext(ctx context.Context) NodeConfigurationPtrOutput {
	return o.ApplyT(func(v NodeConfiguration) *NodeConfiguration {
		return &v
	}).(NodeConfigurationPtrOutput)
}
func (o NodeConfigurationOutput) BootCmd() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NodeConfiguration) []string { return v.BootCmd }).(pulumi.StringArrayOutput)
}

func (o NodeConfigurationOutput) Datasources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NodeConfiguration) []string { return v.Datasources }).(pulumi.StringArrayOutput)
}

func (o NodeConfigurationOutput) Hostname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NodeConfiguration) *string { return v.Hostname }).(pulumi.StringPtrOutput)
}

func (o NodeConfigurationOutput) InitCmd() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NodeConfiguration) []string { return v.InitCmd }).(pulumi.StringArrayOutput)
}

func (o NodeConfigurationOutput) K3OS() K3OSPtrOutput {
	return o.ApplyT(func(v NodeConfiguration) *K3OS { return v.K3OS }).(K3OSPtrOutput)
}

func (o NodeConfigurationOutput) Modules() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NodeConfiguration) []string { return v.Modules }).(pulumi.StringArrayOutput)
}

func (o NodeConfigurationOutput) RunCmd() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NodeConfiguration) []string { return v.RunCmd }).(pulumi.StringArrayOutput)
}

func (o NodeConfigurationOutput) SshAuthorizerKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NodeConfiguration) []string { return v.SshAuthorizerKeys }).(pulumi.StringArrayOutput)
}

func (o NodeConfigurationOutput) Sysctls() pulumi.StringMapOutput {
	return o.ApplyT(func(v NodeConfiguration) map[string]string { return v.Sysctls }).(pulumi.StringMapOutput)
}

func (o NodeConfigurationOutput) WriteFiles() CloudInitFilesPtrOutput {
	return o.ApplyT(func(v NodeConfiguration) *CloudInitFiles { return v.WriteFiles }).(CloudInitFilesPtrOutput)
}

type NodeConfigurationPtrOutput struct{ *pulumi.OutputState }

func (NodeConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NodeConfiguration)(nil)).Elem()
}

func (o NodeConfigurationPtrOutput) ToNodeConfigurationPtrOutput() NodeConfigurationPtrOutput {
	return o
}

func (o NodeConfigurationPtrOutput) ToNodeConfigurationPtrOutputWithContext(ctx context.Context) NodeConfigurationPtrOutput {
	return o
}

func (o NodeConfigurationPtrOutput) Elem() NodeConfigurationOutput {
	return o.ApplyT(func(v *NodeConfiguration) NodeConfiguration { return *v }).(NodeConfigurationOutput)
}

func (o NodeConfigurationPtrOutput) BootCmd() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NodeConfiguration) []string {
		if v == nil {
			return nil
		}
		return v.BootCmd
	}).(pulumi.StringArrayOutput)
}

func (o NodeConfigurationPtrOutput) Datasources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NodeConfiguration) []string {
		if v == nil {
			return nil
		}
		return v.Datasources
	}).(pulumi.StringArrayOutput)
}

func (o NodeConfigurationPtrOutput) Hostname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NodeConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.Hostname
	}).(pulumi.StringPtrOutput)
}

func (o NodeConfigurationPtrOutput) InitCmd() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NodeConfiguration) []string {
		if v == nil {
			return nil
		}
		return v.InitCmd
	}).(pulumi.StringArrayOutput)
}

func (o NodeConfigurationPtrOutput) K3OS() K3OSPtrOutput {
	return o.ApplyT(func(v *NodeConfiguration) *K3OS {
		if v == nil {
			return nil
		}
		return v.K3OS
	}).(K3OSPtrOutput)
}

func (o NodeConfigurationPtrOutput) Modules() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NodeConfiguration) []string {
		if v == nil {
			return nil
		}
		return v.Modules
	}).(pulumi.StringArrayOutput)
}

func (o NodeConfigurationPtrOutput) RunCmd() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NodeConfiguration) []string {
		if v == nil {
			return nil
		}
		return v.RunCmd
	}).(pulumi.StringArrayOutput)
}

func (o NodeConfigurationPtrOutput) SshAuthorizerKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NodeConfiguration) []string {
		if v == nil {
			return nil
		}
		return v.SshAuthorizerKeys
	}).(pulumi.StringArrayOutput)
}

func (o NodeConfigurationPtrOutput) Sysctls() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NodeConfiguration) map[string]string {
		if v == nil {
			return nil
		}
		return v.Sysctls
	}).(pulumi.StringMapOutput)
}

func (o NodeConfigurationPtrOutput) WriteFiles() CloudInitFilesPtrOutput {
	return o.ApplyT(func(v *NodeConfiguration) *CloudInitFiles {
		if v == nil {
			return nil
		}
		return v.WriteFiles
	}).(CloudInitFilesPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(CloudInitFilesOutput{})
	pulumi.RegisterOutputType(CloudInitFilesPtrOutput{})
	pulumi.RegisterOutputType(ConnectionOutput{})
	pulumi.RegisterOutputType(ConnectionPtrOutput{})
	pulumi.RegisterOutputType(K3OSOutput{})
	pulumi.RegisterOutputType(K3OSPtrOutput{})
	pulumi.RegisterOutputType(NodeConfigurationOutput{})
	pulumi.RegisterOutputType(NodeConfigurationPtrOutput{})
}