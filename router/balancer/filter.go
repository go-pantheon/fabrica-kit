package balancer

import (
	"context"

	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/selector"
	"github.com/go-pantheon/fabrica-kit/profile"
	"github.com/go-pantheon/fabrica-kit/xcontext"
)

func NewFilter() selector.NodeFilter {
	return func(ctx context.Context, nodes []selector.Node) []selector.Node {
		newNodes := make([]selector.Node, 0, len(nodes))
		for _, n := range nodes {
			if n.Metadata()[profile.COLOR] == getColorFromCtx(ctx) {
				newNodes = append(newNodes, n)
			}
		}
		return newNodes
	}
}

func getColorFromCtx(ctx context.Context) string {
	if md, ok := metadata.FromServerContext(ctx); ok {
		return md.Get(xcontext.CtxColor)
	}
	return profile.Color()
}
