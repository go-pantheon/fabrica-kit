package redis

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

// Metrics holds Redis-specific metrics
type Metrics struct {
	connectionCount  metric.Int64UpDownCounter
	commandCount     metric.Int64Counter
	commandDuration  metric.Float64Histogram
	memoryUsage      metric.Int64Gauge
	keyspaceHits     metric.Int64Counter
	keyspaceMisses   metric.Int64Counter
	connectedClients metric.Int64Gauge
	errorCount       metric.Int64Counter
}

// MetricsConfig holds configuration for Redis metrics
type MetricsConfig struct {
	MeterProvider metric.MeterProvider
	MeterName     string
	Namespace     string
}

// DefaultMetricsConfig returns default metrics configuration
func DefaultMetricsConfig() *MetricsConfig {
	return &MetricsConfig{
		MeterProvider: otel.GetMeterProvider(),
		MeterName:     "redis-client",
		Namespace:     "redis",
	}
}

// NewMetrics creates Redis-specific metrics
func NewMetrics(config *MetricsConfig) (*Metrics, error) {
	if config == nil {
		config = DefaultMetricsConfig()
	}

	meter := config.MeterProvider.Meter(config.MeterName)

	connectionCount, err := meter.Int64UpDownCounter(
		config.Namespace+".connections.active",
		metric.WithDescription("Number of active Redis connections"),
		metric.WithUnit("connections"),
	)
	if err != nil {
		return nil, err
	}

	commandCount, err := meter.Int64Counter(
		config.Namespace+".commands.total",
		metric.WithDescription("Total number of Redis commands executed"),
		metric.WithUnit("commands"),
	)
	if err != nil {
		return nil, err
	}

	commandDuration, err := meter.Float64Histogram(
		config.Namespace+".commands.duration",
		metric.WithDescription("Duration of Redis commands"),
		metric.WithUnit("ms"),
		metric.WithExplicitBucketBoundaries(0.1, 0.5, 1.0, 2.5, 5.0, 10.0, 25.0, 50.0, 100.0, 250.0, 500.0, 1000.0),
	)
	if err != nil {
		return nil, err
	}

	memoryUsage, err := meter.Int64Gauge(
		config.Namespace+".memory.used",
		metric.WithDescription("Redis memory usage in bytes"),
		metric.WithUnit("bytes"),
	)
	if err != nil {
		return nil, err
	}

	keyspaceHits, err := meter.Int64Counter(
		config.Namespace+".keyspace.hits",
		metric.WithDescription("Number of successful key lookups"),
		metric.WithUnit("hits"),
	)
	if err != nil {
		return nil, err
	}

	keyspaceMisses, err := meter.Int64Counter(
		config.Namespace+".keyspace.misses",
		metric.WithDescription("Number of failed key lookups"),
		metric.WithUnit("misses"),
	)
	if err != nil {
		return nil, err
	}

	connectedClients, err := meter.Int64Gauge(
		config.Namespace+".clients.connected",
		metric.WithDescription("Number of connected clients"),
		metric.WithUnit("clients"),
	)
	if err != nil {
		return nil, err
	}

	errorCount, err := meter.Int64Counter(
		config.Namespace+".errors.total",
		metric.WithDescription("Total number of Redis errors"),
		metric.WithUnit("errors"),
	)
	if err != nil {
		return nil, err
	}

	return &Metrics{
		connectionCount:  connectionCount,
		commandCount:     commandCount,
		commandDuration:  commandDuration,
		memoryUsage:      memoryUsage,
		keyspaceHits:     keyspaceHits,
		keyspaceMisses:   keyspaceMisses,
		connectedClients: connectedClients,
		errorCount:       errorCount,
	}, nil
}

// RecordConnection records a connection event
func (m *Metrics) RecordConnection(ctx context.Context, addr string, delta int64) {
	m.connectionCount.Add(ctx, delta, metric.WithAttributes(
		attribute.String("redis.address", addr),
	))
}

// RecordClusterConnection records a cluster connection event
func (m *Metrics) RecordClusterConnection(ctx context.Context, addrs []string, delta int64) {
	m.connectionCount.Add(ctx, delta, metric.WithAttributes(
		attribute.StringSlice("redis.addresses", addrs),
		attribute.String("redis.type", "cluster"),
	))
}

// RecordCommand records a command execution
func (m *Metrics) RecordCommand(ctx context.Context, command string, duration float64) {
	attrs := metric.WithAttributes(
		attribute.String("redis.command", command),
	)

	m.commandCount.Add(ctx, 1, attrs)
	m.commandDuration.Record(ctx, duration, attrs)
}

// RecordError records an error
func (m *Metrics) RecordError(ctx context.Context, command string, errorType string) {
	m.errorCount.Add(ctx, 1, metric.WithAttributes(
		attribute.String("redis.command", command),
		attribute.String("redis.error.type", errorType),
	))
}

// RecordKeyspaceHit records a keyspace hit
func (m *Metrics) RecordKeyspaceHit(ctx context.Context) {
	m.keyspaceHits.Add(ctx, 1)
}

// RecordKeyspaceMiss records a keyspace miss
func (m *Metrics) RecordKeyspaceMiss(ctx context.Context) {
	m.keyspaceMisses.Add(ctx, 1)
}

// RecordMemoryUsage records current memory usage
func (m *Metrics) RecordMemoryUsage(ctx context.Context, bytes int64) {
	m.memoryUsage.Record(ctx, bytes)
}

// RecordConnectedClients records the number of connected clients
func (m *Metrics) RecordConnectedClients(ctx context.Context, count int64) {
	m.connectedClients.Record(ctx, count)
}
