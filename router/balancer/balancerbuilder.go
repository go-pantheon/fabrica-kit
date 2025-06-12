package balancer

import (
	"github.com/go-kratos/kratos/v2/selector"
	"github.com/go-kratos/kratos/v2/selector/node/direct"
	"github.com/go-pantheon/fabrica-kit/router/routetable"
)

// Option is a function that configures the selector options.
type Option func(o *options)

type options struct {
	balancerType Type
	routeTable   routetable.RouteTable
}

// WithRouteTable sets the route table for the balancer.
func WithRouteTable(rt routetable.RouteTable) Option {
	return func(o *options) {
		o.routeTable = rt
	}
}

// WithBalancerType sets the balancer type.
func WithBalancerType(balancerType Type) Option {
	return func(o *options) {
		o.balancerType = balancerType
	}
}

var _ selector.BalancerBuilder = (*balancerBuilder)(nil)

// balancerBuilder is a selector builder for creating weighted round-robin balancers.
type balancerBuilder struct {
	balancerType Type
	routeTable   routetable.RouteTable
}

// newBalancerBuilder returns a selector builder with wrr balancer
func newBalancerBuilder(opts ...Option) selector.Builder {
	var option options
	for _, opt := range opts {
		opt(&option)
	}

	return &selector.DefaultBuilder{
		Balancer: &balancerBuilder{
			balancerType: option.balancerType,
			routeTable:   option.routeTable,
		},
		Node: &direct.Builder{},
	}
}

// Build creates a new balancer instance.
func (b *balancerBuilder) Build() selector.Balancer {
	return newWeightBalancer(b.balancerType, b.routeTable)
}
