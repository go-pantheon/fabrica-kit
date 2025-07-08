package traceredis

import (
	"context"

	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// TracingConfig holds configuration for Redis tracing
type TracingConfig struct {
	TracerProvider trace.TracerProvider
	ServiceName    string
	ServiceVersion string
}

// DefaultTracingConfig returns default tracing configuration
func DefaultTracingConfig() *TracingConfig {
	return &TracingConfig{
		TracerProvider: otel.GetTracerProvider(),
		ServiceName:    "redis-client",
		ServiceVersion: "1.0.0",
	}
}

// WithTracing adds OpenTelemetry tracing to a Cacheable interface from fabrica-util
func WithTracing(rdb redis.UniversalClient, config *TracingConfig) error {
	if config == nil {
		config = DefaultTracingConfig()
	}

	return redisotel.InstrumentTracing(
		rdb.(*redis.Client),
		redisotel.WithTracerProvider(config.TracerProvider),
	)
}

// StartConnectionSpan creates a span for Redis connection operations
func StartConnectionSpan(ctx context.Context, operation string, addr string, db int) (context.Context, trace.Span) {
	tracer := otel.Tracer("redis-client")
	ctx, span := tracer.Start(ctx, "redis."+operation)

	span.SetAttributes(
		attribute.String("redis.address", addr),
		attribute.String("redis.operation", operation),
		attribute.Int("redis.db", db),
	)

	return ctx, span
}

// StartClusterConnectionSpan creates a span for Redis cluster connection operations
func StartClusterConnectionSpan(ctx context.Context, operation string, addrs []string) (context.Context, trace.Span) {
	tracer := otel.Tracer("redis-client")
	ctx, span := tracer.Start(ctx, "redis.cluster."+operation)

	span.SetAttributes(
		attribute.StringSlice("redis.addresses", addrs),
		attribute.String("redis.operation", operation),
		attribute.String("redis.type", "cluster"),
	)

	return ctx, span
}

// SetConnectionSuccess marks a connection span as successful
func SetConnectionSuccess(span trace.Span) {
	span.SetAttributes(attribute.Bool("redis.connected", true))
}

// SetConnectionError records an error on a connection span
func SetConnectionError(span trace.Span, err error) {
	span.RecordError(err)
	span.SetAttributes(attribute.Bool("redis.connected", false))
}
