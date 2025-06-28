package metadata

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
	"google.golang.org/grpc/metadata"
)

type metadataKey struct{}

// Server is a server middleware that retrieves metadata from the incoming context
// and stores it in the context for later use.
func Server() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (any, error) {
			if md, ok := metadata.FromIncomingContext(ctx); ok {
				ctx = context.WithValue(ctx, metadataKey{}, md.Copy())
			}
			return handler(ctx, req)
		}
	}
}

// Client is a client middleware that retrieves metadata from the context
// and appends it to the outgoing gRPC context.
func Client() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (any, error) {
			if md, ok := FromContext(ctx); ok {
				existingMd, ok := metadata.FromOutgoingContext(ctx)
				if ok {
					ctx = metadata.NewOutgoingContext(ctx, metadata.Join(existingMd, md))
				} else {
					ctx = metadata.NewOutgoingContext(ctx, md)
				}
			}
			return handler(ctx, req)
		}
	}
}

// FromContext returns the metadata from the given context.
func FromContext(ctx context.Context) (metadata.MD, bool) {
	md, ok := ctx.Value(metadataKey{}).(metadata.MD)
	return md, ok
}
