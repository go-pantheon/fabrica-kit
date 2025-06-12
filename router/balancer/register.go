package balancer

import (
	"sync/atomic"

	"github.com/go-kratos/kratos/v2/selector"
	"github.com/go-pantheon/fabrica-kit/router/routetable"
	"github.com/go-pantheon/fabrica-util/errors"
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
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
	// readerBalancerBuilderRegistered indicates whether the reader balancer has been registered.
	readerBalancerBuilderRegistered atomic.Bool
	// masterBalancerBuilderRegistered indicates whether the master balancer has been registered.
	masterBalancerBuilderRegistered atomic.Bool
)

// RegisterMasterBalancer registers a balancer for master nodes.
// It uses the provided route table for routing decisions.
func RegisterMasterBalancer(rt routetable.RouteTable) error {
	if masterBalancerBuilderRegistered.Load() {
		return errors.New("master balancer already registered")
	}

	t := TypeMaster
	registerBalancerBuilder(t, newBalancerBuilder(WithBalancerType(t), WithRouteTable(rt)))
	masterBalancerBuilderRegistered.Store(true)

	return nil
}

// RegisterReaderBalancer registers a balancer for reader nodes.
// It uses the provided route table for routing decisions.
func RegisterReaderBalancer(rt routetable.RouteTable) error {
	if readerBalancerBuilderRegistered.Load() {
		return errors.New("reader balancer already registered")
	}

	t := TypeReader
	registerBalancerBuilder(t, newBalancerBuilder(WithBalancerType(t), WithRouteTable(rt)))
	readerBalancerBuilderRegistered.Store(true)

	return nil
}

func registerBalancerBuilder(balancerType Type, builder selector.Builder) {
	b := base.NewBalancerBuilder(
		string(balancerType),
		newPickerBuilder(builder),
		base.Config{HealthCheck: true},
	)
	balancer.Register(b)
}
