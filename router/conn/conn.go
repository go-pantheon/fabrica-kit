package conn

import (
	"context"
	"fmt"

	"github.com/luffy050596/vulcan-pkg-app/router/routetable"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type Conn struct {
	grpc.ClientConnInterface
}

func NewConn(serviceName string, balancerType balancer.BalancerType, logger log.Logger, rt routetable.RouteTable, r registry.Discovery) (*Conn, error) {
	conn, err := kgrpc.DialInsecure(
		context.Background(),
		kgrpc.WithEndpoint(fmt.Sprintf("discovery:///%s", serviceName)),
		kgrpc.WithDiscovery(r),
		kgrpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		return nil, errors.Wrapf(err, "create grpc connection failed. app=%s", serviceName)
	}
	return &Conn{ClientConnInterface: conn}, nil
}
