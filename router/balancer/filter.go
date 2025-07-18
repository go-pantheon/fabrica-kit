package balancer

import (
	"context"

	"github.com/go-kratos/kratos/v2/selector"
	"github.com/go-pantheon/fabrica-kit/profile"
	"github.com/go-pantheon/fabrica-kit/xcontext"
)

// NewFilter creates a node filter that filters nodes based on color.
// It returns a selector.NodeFilter that selects nodes matching the color from context.
func NewFilter() selector.NodeFilter {
	return func(ctx context.Context, nodes []selector.Node) []selector.Node {
		newNodes := make([]selector.Node, 0, len(nodes))

		for _, n := range nodes {
			if n.Metadata()[profile.ColorKey] == xcontext.ColorFromOutgoingContext(ctx) {
				newNodes = append(newNodes, n)
			}
		}

		return newNodes
	}
}
