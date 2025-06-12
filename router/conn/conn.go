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
	kgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-pantheon/fabrica-kit/metrics"
	"github.com/go-pantheon/fabrica-kit/router/balancer"
	"github.com/go-pantheon/fabrica-kit/router/routetable"
	"github.com/go-pantheon/fabrica-util/errors"
	"google.golang.org/grpc"
)

// Conn is a wrapper around a gRPC client connection interface.
type Conn struct {
	grpc.ClientConnInterface
}

// NewConn creates a new gRPC client connection with the specified service name, balancer type,
// logger, route table, and discovery mechanism.
// It configures the connection with appropriate middleware and balancer settings.
func NewConn(serviceName string, balancerType balancer.Type, logger log.Logger, rt routetable.RouteTable, r registry.Discovery) (*Conn, error) {
	switch balancerType {
	case balancer.TypeMaster:
		if err := balancer.RegisterMasterBalancer(rt); err != nil {
			return nil, err
		}
	case balancer.TypeReader:
		if err := balancer.RegisterReaderBalancer(rt); err != nil {
			return nil, err
		}
	default:
		return nil, errors.Errorf("invalid balancer type: %s", balancerType)
	}

	conn, err := kgrpc.DialInsecure(
		context.Background(),
		kgrpc.WithEndpoint(fmt.Sprintf("discovery:///%s", serviceName)),
		kgrpc.WithDiscovery(r),
		kgrpc.WithNodeFilter(balancer.NewFilter()),
		kgrpc.WithOptions(
			grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"loadBalancingConfig": [{"%s":{}}]}`, string(balancerType))),
		),
		kgrpc.WithMiddleware(
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
