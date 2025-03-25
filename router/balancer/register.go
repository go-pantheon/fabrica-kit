package balancer

import (
	"sync/atomic"

	"github.com/go-pantheon/fabrica-kit/router/routetable"
)

type BalancerType string

const (
	BalancerTypeMaster BalancerType = "master"
	BalancerTypeViewer BalancerType = "viewer"
)

var (
	ReaderBalancerRegistered atomic.Bool
	MasterBalancerRegistered atomic.Bool
)

// RegisterBalancer Register a balancer for master
// return the balancer name
func RegisterMasterBalancer(rt routetable.RouteTable) {
	t := BalancerTypeMaster
	registerBalancer(t, NewBuilder(WithBalancerType(t), WithRouteTable(rt)))
	MasterBalancerRegistered.Store(true)
}

// RegisterBalancer Register a balancer for reader
// return the balancer name
func RegisterReaderBalancer(rt routetable.RouteTable) {
	t := BalancerTypeViewer
	registerBalancer(t, NewBuilder(WithBalancerType(t), WithRouteTable(rt)))
	ReaderBalancerRegistered.Store(true)
}
