package balancer

import (
	"sync/atomic"

	"github.com/go-pantheon/fabrica-kit/router/routetable"
)

// Type represents the type of load balancer.
type Type string

const (
	// TypeMaster is the balancer type for master nodes.
	TypeMaster Type = "master"
	// TypeReader is the balancer type for reader nodes.
	TypeReader Type = "reader"
)

var (
	// ReaderBalancerRegistered indicates whether the reader balancer has been registered.
	ReaderBalancerRegistered atomic.Bool
	// MasterBalancerRegistered indicates whether the master balancer has been registered.
	MasterBalancerRegistered atomic.Bool
)

// RegisterMasterBalancer registers a balancer for master nodes.
// It uses the provided route table for routing decisions.
func RegisterMasterBalancer(rt routetable.RouteTable) {
	t := TypeMaster
	registerBalancer(t, NewBuilder(WithBalancerType(t), WithRouteTable(rt)))
	MasterBalancerRegistered.Store(true)
}

// RegisterReaderBalancer registers a balancer for reader nodes.
// It uses the provided route table for routing decisions.
func RegisterReaderBalancer(rt routetable.RouteTable) {
	t := TypeReader
	registerBalancer(t, NewBuilder(WithBalancerType(t), WithRouteTable(rt)))
	ReaderBalancerRegistered.Store(true)
}
