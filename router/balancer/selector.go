package balancer

import (
	"context"
	"strconv"
	"sync"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/selector"
	"github.com/go-kratos/kratos/v2/selector/node/direct"
	"github.com/go-pantheon/fabrica-kit/router/routetable"
	"github.com/go-pantheon/fabrica-kit/xcontext"
	"github.com/pkg/errors"
)

// New random a selector.
func New(opts ...Option) selector.Selector {
	return NewBuilder(opts...).Build()
}

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

// Builder is a selector builder for creating weighted round-robin balancers.
type Builder struct {
	balancerType Type
	routeTable   routetable.RouteTable
}

// NewBuilder returns a selector builder with wrr balancer
func NewBuilder(opts ...Option) selector.Builder {
	var option options
	for _, opt := range opts {
		opt(&option)
	}

	return &selector.DefaultBuilder{
		Balancer: &Builder{
			balancerType: option.balancerType,
			routeTable:   option.routeTable,
		},
		Node: &direct.Builder{},
	}
}

// Build creates a new balancer instance.
func (b *Builder) Build() selector.Balancer {
	return &Balancer{
		balancerType:  b.balancerType,
		currentWeight: make(map[string]float64),
		routeTable:    b.routeTable,
	}
}

var _ selector.Balancer = (*Balancer)(nil)

// Balancer is a weighted round-robin load balancer that supports route tables.
type Balancer struct {
	balancerType  Type
	mu            sync.Mutex
	currentWeight map[string]float64
	routeTable    routetable.RouteTable
}

// Pick is pick a weighted node
func (p *Balancer) Pick(ctx context.Context, nodes []selector.WeightedNode) (selector.WeightedNode, selector.DoneFunc, error) {
	if len(nodes) == 0 {
		return nil, nil, selector.ErrNoAvailable
	}

	oid, err := getOIDFromCtx(ctx)
	if err != nil {
		return nil, nil, err
	}

	color := getColorFromCtx(ctx)

	// select node by oid from routeTable
	addr, err := p.routeTable.LoadAndExpire(ctx, color, oid)
	if err != nil {
		return nil, nil, err
	}

	for _, node := range nodes {
		if node.Address() == addr {
			return node, nil, nil
		}
	}

	// select a new node by weight from nodes
	// the algorithm is the implement of nginx wrr, copied from https://github.com/go-kratos/kratos/blob/main/selector/wrr/wrr.go
	var (
		totalWeight  float64
		selected     selector.WeightedNode
		selectWeight float64
	)

	p.mu.Lock()

	for _, node := range nodes {
		totalWeight += node.Weight()
		cwt := p.currentWeight[node.Address()]
		cwt += node.Weight()
		p.currentWeight[node.Address()] = cwt

		if selected == nil || selectWeight < cwt {
			selectWeight = cwt
			selected = node
		}
	}

	p.currentWeight[selected.Address()] = selectWeight - totalWeight
	p.mu.Unlock()

	d := selected.Pick()

	// the select action is done, return the selected node if the balancer type is not master
	if p.balancerType != TypeMaster {
		return selected, d, nil
	}

	// update route table if the balancer type is master
	// the route table may be set by other connections at the same time, so we need to judge it with SetNx before setting
	ok, addr, err := p.routeTable.SetNx(ctx, color, oid, selected.Address())
	if err != nil {
		return nil, nil, err
	}

	if ok {
		// the route table is set by this balancer
		return selected, d, nil
	}

	log.Warnf("routeTable is set by other balancers. oid=%d color=%s oldConn=%s newConn=%s", oid, color, addr, selected.Address())

	for _, node := range nodes {
		if node.Address() == addr {
			return node, nil, nil
		}
	}

	return nil, nil, errors.Errorf("the existed connection in routeTable is not found. oid=%d color=%s oldConn=%s", oid, color, addr)
}

func getOIDFromCtx(ctx context.Context) (oid int64, err error) {
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		err = errors.Errorf("metadata not in context")

		return
	}

	if oid, err = strconv.ParseInt(md.Get(xcontext.CtxOID), 10, 64); err != nil {
		err = errors.Wrapf(err, "oid not int64")
	}

	return
}
