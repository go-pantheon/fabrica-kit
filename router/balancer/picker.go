package balancer

import (
	"github.com/go-kratos/kratos/v2/selector"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/metadata"
)

var _ balancer.Picker = (*picker)(nil)

// picker is a grpc picker.
type picker struct {
	selector selector.Selector
}

func newPicker(selector selector.Selector) *picker {
	return &picker{
		selector: selector,
	}
}

// Pick pick instances.
func (p *picker) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	var filters []selector.NodeFilter

	if tr, ok := transport.FromClientContext(info.Ctx); ok {
		if gtr, ok := tr.(*grpc.Transport); ok {
			filters = gtr.NodeFilters()
		}
	}

	n, done, err := p.selector.Select(info.Ctx, selector.WithNodeFilter(filters...))
	if err != nil {
		return balancer.PickResult{}, err
	}

	return balancer.PickResult{
		SubConn: n.(*grpcNode).subConn,
		Done: func(di balancer.DoneInfo) {
			done(info.Ctx, selector.DoneInfo{
				Err:           di.Err,
				BytesSent:     di.BytesSent,
				BytesReceived: di.BytesReceived,
				ReplyMD:       Trailer(di.Trailer),
			})
		},
	}, nil
}

// Trailer is a grpc trailer MD.
type Trailer metadata.MD

// Get get a grpc trailer value.
func (t Trailer) Get(k string) string {
	v := metadata.MD(t).Get(k)
	if len(v) > 0 {
		return v[0]
	}

	return ""
}
