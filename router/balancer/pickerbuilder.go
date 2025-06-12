// Package balancer provides gRPC load balancing functionality
// for service discovery and routing in distributed systems.
package balancer

import (
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/selector"
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
)

var _ base.PickerBuilder = (*pickerBuilder)(nil)

type pickerBuilder struct {
	builder selector.Builder
}

func newPickerBuilder(builder selector.Builder) *pickerBuilder {
	return &pickerBuilder{
		builder: builder,
	}
}

// Build creates a grpc Picker.
func (b *pickerBuilder) Build(info base.PickerBuildInfo) balancer.Picker {
	if len(info.ReadySCs) == 0 {
		// Block the RPC until a new picker is available via UpdateState().
		return base.NewErrPicker(balancer.ErrNoSubConnAvailable)
	}

	nodes := make([]selector.Node, 0, len(info.ReadySCs))

	for conn, info := range info.ReadySCs {
		ins, _ := info.Address.Attributes.Value("rawServiceInstance").(*registry.ServiceInstance)
		nodes = append(nodes, newGrpcNode(selector.NewNode("grpc", info.Address.Addr, ins), conn))
	}

	p := newPicker(b.builder.Build())
	p.selector.Apply(nodes)

	return p
}

var _ selector.Node = (*grpcNode)(nil)

type grpcNode struct {
	selector.Node

	subConn balancer.SubConn
}

func newGrpcNode(node selector.Node, subConn balancer.SubConn) *grpcNode {
	return &grpcNode{
		Node:    node,
		subConn: subConn,
	}
}
