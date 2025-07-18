// Package conn provides utilities for managing gRPC client connections
// with load balancing, service discovery, and middleware support.
package conn

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-pantheon/fabrica-kit/metrics"
	"github.com/go-pantheon/fabrica-kit/router/balancer"
	"github.com/go-pantheon/fabrica-kit/router/routetable"
	"github.com/go-pantheon/fabrica-util/errors"
	grpcgo "google.golang.org/grpc"
)

// Conn is a wrapper around a gRPC client connection interface.
type Conn struct {
	grpcgo.ClientConnInterface
}

// NewConn creates a new gRPC client connection with the specified service name, balancer type,
// logger, route table, and discovery mechanism.
// It configures the connection with appropriate middleware and balancer settings.
func NewConn(serviceName string, balancerType balancer.Type, logger log.Logger, rt routetable.ReadOnlyRouteTable, r registry.Discovery) (*Conn, error) {
	switch balancerType {
	case balancer.TypeMaster:
		mrt, ok := rt.(routetable.MasterRouteTable)
		if !ok {
			return nil, errors.Errorf("route table is not a master route table")
		}

		balancer.RegisterMasterBalancer(mrt)
	case balancer.TypeReader:
		balancer.RegisterReadOnlyBalancer(rt)
	default:
		return nil, errors.Errorf("invalid balancer type: %s", balancerType)
	}

	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(fmt.Sprintf("discovery:///%s", serviceName)),
		grpc.WithDiscovery(r),
		grpc.WithNodeFilter(balancer.NewFilter()),
		grpc.WithOptions(
			grpcgo.WithDefaultServiceConfig(fmt.Sprintf(`{"loadBalancingConfig": [{"%s":{}}]}`, string(balancerType))),
		),
		grpc.WithMiddleware(
			recovery.Recovery(),
			metadata.Client(),
			tracing.Client(),
			metrics.Client(),
			logging.Client(logger),
		),
	)
	if err != nil {
		return nil, errors.Wrapf(err, "create grpc connection failed. app=%s", serviceName)
	}

	return &Conn{ClientConnInterface: conn}, nil
}
