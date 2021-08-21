// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package k3os

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Node struct {
	pulumi.CustomResourceState
}

// NewNode registers a new resource with the given unique name, arguments, and options.
func NewNode(ctx *pulumi.Context,
	name string, args *NodeArgs, opts ...pulumi.ResourceOption) (*Node, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Connection == nil {
		return nil, errors.New("invalid value for required argument 'Connection'")
	}
	if args.NodeConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'NodeConfiguration'")
	}
	var resource Node
	err := ctx.RegisterResource("k3os:index:Node", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNode gets an existing Node resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNode(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NodeState, opts ...pulumi.ResourceOption) (*Node, error) {
	var resource Node
	err := ctx.ReadResource("k3os:index:Node", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Node resources.
type nodeState struct {
}

type NodeState struct {
}

func (NodeState) ElementType() reflect.Type {
	return reflect.TypeOf((*nodeState)(nil)).Elem()
}

type nodeArgs struct {
	Connection        Connection        `pulumi:"connection"`
	NodeConfiguration NodeConfiguration `pulumi:"nodeConfiguration"`
}

// The set of arguments for constructing a Node resource.
type NodeArgs struct {
	Connection        ConnectionInput
	NodeConfiguration NodeConfigurationInput
}

func (NodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nodeArgs)(nil)).Elem()
}

type NodeInput interface {
	pulumi.Input

	ToNodeOutput() NodeOutput
	ToNodeOutputWithContext(ctx context.Context) NodeOutput
}

func (*Node) ElementType() reflect.Type {
	return reflect.TypeOf((*Node)(nil))
}

func (i *Node) ToNodeOutput() NodeOutput {
	return i.ToNodeOutputWithContext(context.Background())
}

func (i *Node) ToNodeOutputWithContext(ctx context.Context) NodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeOutput)
}

type NodeOutput struct {
	*pulumi.OutputState
}

func (NodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Node)(nil))
}

func (o NodeOutput) ToNodeOutput() NodeOutput {
	return o
}

func (o NodeOutput) ToNodeOutputWithContext(ctx context.Context) NodeOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NodeOutput{})
}
