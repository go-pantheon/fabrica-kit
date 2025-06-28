package balancer

import (
	"context"
	"strconv"
	"sync"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/selector"
	"github.com/go-pantheon/fabrica-kit/router/routetable"
	"github.com/go-pantheon/fabrica-kit/xcontext"
	"github.com/go-pantheon/fabrica-kit/xerrors"
	"github.com/go-pantheon/fabrica-util/errors"
)

var _ selector.Balancer = (*weightBalancer)(nil)

// weightBalancer is a weighted round-robin load balancer that supports route tables.
type weightBalancer struct {
	mu sync.Mutex

	balancerType  Type
	currentWeight map[string]float64
	routeTable    routetable.ReadOnlyRouteTable
}

func newWeightBalancer(balancerType Type, routeTable routetable.ReadOnlyRouteTable) selector.Balancer {
	return &weightBalancer{
		balancerType:  balancerType,
		currentWeight: make(map[string]float64),
		routeTable:    routeTable,
	}
}

var emptyDoneFunc = func(ctx context.Context, di selector.DoneInfo) {}

// Pick is pick a weighted node
func (p *weightBalancer) Pick(ctx context.Context, nodes []selector.WeightedNode) (selector.WeightedNode, selector.DoneFunc, error) {
	if len(nodes) == 0 {
		return nil, nil, selector.ErrNoAvailable
	}

	oid, err := getOIDFromCtx(ctx)
	if err != nil {
		return nil, nil, err
	}

	color := getColorFromCtx(ctx)

	// select node by oid from routeTable
	addr, err := p.routeTable.Get(ctx, color, oid)
	if err != nil && !errors.Is(err, xerrors.ErrRouteTableNotFound) {
		return nil, nil, err
	}

	for _, node := range nodes {
		if node.Address() == addr {
			return node, emptyDoneFunc, nil
		}
	}

	selected := p.weightSelect(nodes)
	if selected == nil {
		return nil, nil, errors.New("the selected node is nil")
	}

	// the select action is done, return the selected node if the balancer type is not master
	if p.balancerType != TypeMaster {
		return selected, selected.Pick(), nil
	}

	mrt, ok := p.routeTable.(routetable.MasterRouteTable)
	if !ok {
		return nil, nil, errors.New("the route table is not a RouteTable")
	}

	// update route table if the balancer type is master
	// the route table may be set by other connections at the same time, so we need to judge it with SetNx before setting
	ok, addr, err = mrt.SetNxOrGet(ctx, color, oid, selected.Address())
	if err != nil {
		return nil, nil, err
	}

	if ok {
		// the route table is set by this balancer
		return selected, selected.Pick(), nil
	}

	log.Warnf("routeTable is set by other balancers. oid=%d color=%s old-addr=%s new-addr=%s", oid, color, addr, selected.Address())

	for _, node := range nodes {
		if node.Address() == addr {
			return node, node.Pick(), nil
		}
	}

	return nil, nil, errors.Errorf("the existed address in routeTable is not found. addr=%s", addr)
}

// weightSelect select a new node by weight from nodes
// the algorithm is the implement of nginx wrr, copied from https://github.com/go-kratos/kratos/blob/main/selector/wrr/wrr.go
func (p *weightBalancer) weightSelect(nodes []selector.WeightedNode) selector.WeightedNode {
	var (
		totalWeight  float64
		selected     selector.WeightedNode
		selectWeight float64
	)

	p.mu.Lock()
	defer p.mu.Unlock()

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

	return selected
}

func getOIDFromCtx(ctx context.Context) (int64, error) {
	md, ok := metadata.FromClientContext(ctx)
	if !ok {
		return 0, errors.New("metadata is not in context")
	}

	oid, err := strconv.ParseInt(md.Get(xcontext.CtxOID), 10, 64)
	if err != nil {
		return 0, errors.Wrapf(err, "oid is not int64")
	}

	return oid, nil
}
