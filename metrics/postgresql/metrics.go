package postgresql

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-pantheon/fabrica-util/errors"
	"github.com/go-pantheon/fabrica-util/xsync"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

// Metrics holds all PostgreSQL-related metrics
type Metrics struct {
	meter metric.Meter

	// Connection pool metrics
	poolActiveConnections  metric.Int64UpDownCounter
	poolIdleConnections    metric.Int64UpDownCounter
	poolOpenConnections    metric.Int64UpDownCounter
	poolMaxOpenConnections metric.Int64UpDownCounter
	poolWaitCount          metric.Int64Counter
	poolWaitDuration       metric.Float64Histogram
	poolIdleClosed         metric.Int64Counter
	poolLifetimeClosed     metric.Int64Counter

	// Query metrics
	queryDuration metric.Float64Histogram
	queryCounter  metric.Int64Counter
	queryErrors   metric.Int64Counter

	// Health check metrics
	healthCheckDuration metric.Float64Histogram
	healthCheckCounter  metric.Int64Counter
	healthCheckErrors   metric.Int64Counter

	// Common attributes
	dbName   string
	dbSystem string
}

// MetricsConfig holds configuration for PostgreSQL metrics
type MetricsConfig struct {
	ServiceName         string
	DBName              string
	DBSystem            string
	MeterProvider       metric.MeterProvider
	StatsReportInterval time.Duration
	EnableQueryMetrics  bool
	EnablePoolMetrics   bool
	EnableHealthMetrics bool
}

// DefaultMetricsConfig returns default configuration
func DefaultMetricsConfig(serviceName, dbName string) *MetricsConfig {
	if serviceName == "" {
		serviceName = "postgres-metrics"
	}

	return &MetricsConfig{
		ServiceName:         serviceName,
		DBName:              dbName,
		DBSystem:            "postgresql",
		MeterProvider:       otel.GetMeterProvider(),
		StatsReportInterval: 30 * time.Second,
		EnableQueryMetrics:  true,
		EnablePoolMetrics:   true,
		EnableHealthMetrics: true,
	}
}

// NewMetrics creates a new PostgreSQL metrics collector
func NewMetrics(config *MetricsConfig) (*Metrics, error) {
	if config == nil {
		return nil, errors.New("config cannot be nil")
	}

	meter := config.MeterProvider.Meter(config.ServiceName)

	m := &Metrics{
		meter:    meter,
		dbName:   config.DBName,
		dbSystem: config.DBSystem,
	}

	var err error

	// Initialize connection pool metrics
	if config.EnablePoolMetrics {
		if err = m.initPoolMetrics(); err != nil {
			return nil, errors.Wrap(err, "failed to initialize pool metrics")
		}
	}

	// Initialize query metrics
	if config.EnableQueryMetrics {
		if err = m.initQueryMetrics(); err != nil {
			return nil, errors.Wrap(err, "failed to initialize query metrics")
		}
	}

	// Initialize health check metrics
	if config.EnableHealthMetrics {
		if err = m.initHealthMetrics(); err != nil {
			return nil, errors.Wrap(err, "failed to initialize health metrics")
		}
	}

	return m, nil
}

// initPoolMetrics initializes connection pool metrics
func (m *Metrics) initPoolMetrics() error {
	var err error

	m.poolActiveConnections, err = m.meter.Int64UpDownCounter(
		"db.postgres.connections.active",
		metric.WithDescription("Number of active database connections"),
	)
	if err != nil {
		return errors.Wrap(err, "failed to initialize pool active connections")
	}

	m.poolIdleConnections, err = m.meter.Int64UpDownCounter(
		"db.postgres.connections.idle",
		metric.WithDescription("Number of idle database connections"),
	)
	if err != nil {
		return errors.Wrap(err, "failed to initialize pool idle connections")
	}

	m.poolOpenConnections, err = m.meter.Int64UpDownCounter(
		"db.postgres.connections.open",
		metric.WithDescription("Number of open database connections"),
	)
	if err != nil {
		return errors.Wrap(err, "failed to initialize pool open connections")
	}

	m.poolMaxOpenConnections, err = m.meter.Int64UpDownCounter(
		"db.postgres.connections.max_open",
		metric.WithDescription("Maximum number of open database connections allowed"),
	)
	if err != nil {
		return errors.Wrap(err, "failed to initialize pool max open connections")
	}

	m.poolWaitCount, err = m.meter.Int64Counter(
		"db.postgres.connections.wait_count",
		metric.WithDescription("Total number of connections waited for"),
	)
	if err != nil {
		return errors.Wrap(err, "failed to initialize pool wait count")
	}

	m.poolWaitDuration, err = m.meter.Float64Histogram(
		"db.postgres.connections.wait_duration",
		metric.WithDescription("Time blocked waiting for a new connection"),
		metric.WithUnit("s"),
	)
	if err != nil {
		return errors.Wrap(err, "failed to initialize pool wait duration")
	}

	m.poolIdleClosed, err = m.meter.Int64Counter(
		"db.postgres.connections.idle_closed",
		metric.WithDescription("Total number of connections closed due to SetMaxIdleConns"),
	)
	if err != nil {
		return errors.Wrap(err, "failed to initialize pool idle closed")
	}

	m.poolLifetimeClosed, err = m.meter.Int64Counter(
		"db.postgres.connections.lifetime_closed",
		metric.WithDescription("Total number of connections closed due to SetConnMaxLifetime"),
	)
	if err != nil {
		return errors.Wrap(err, "failed to initialize pool lifetime closed")
	}

	return nil
}

// initQueryMetrics initializes query metrics
func (m *Metrics) initQueryMetrics() error {
	var err error

	m.queryDuration, err = m.meter.Float64Histogram(
		"db.postgres.query.duration",
		metric.WithDescription("Duration of database queries"),
		metric.WithUnit("s"),
	)
	if err != nil {
		return errors.Wrap(err, "failed to initialize query duration")
	}

	m.queryCounter, err = m.meter.Int64Counter(
		"db.postgres.query.count",
		metric.WithDescription("Total number of database queries"),
	)
	if err != nil {
		return errors.Wrap(err, "failed to initialize query counter")
	}

	m.queryErrors, err = m.meter.Int64Counter(
		"db.postgres.query.errors",
		metric.WithDescription("Total number of database query errors"),
	)
	if err != nil {
		return errors.Wrap(err, "failed to initialize query errors")
	}

	return nil
}

// initHealthMetrics initializes health check metrics
func (m *Metrics) initHealthMetrics() error {
	var err error

	m.healthCheckDuration, err = m.meter.Float64Histogram(
		"db.postgres.health_check.duration",
		metric.WithDescription("Duration of database health checks"),
		metric.WithUnit("s"),
	)
	if err != nil {
		return errors.Wrap(err, "failed to initialize health check duration")
	}

	m.healthCheckCounter, err = m.meter.Int64Counter(
		"db.postgres.health_check.count",
		metric.WithDescription("Total number of database health checks"),
	)
	if err != nil {
		return errors.Wrap(err, "failed to initialize health check counter")
	}

	m.healthCheckErrors, err = m.meter.Int64Counter(
		"db.postgres.health_check.errors",
		metric.WithDescription("Total number of database health check errors"),
	)
	if err != nil {
		return errors.Wrap(err, "failed to initialize health check errors")
	}

	return nil
}

// commonAttributes returns common attributes for all metrics
func (m *Metrics) commonAttributes() []attribute.KeyValue {
	return []attribute.KeyValue{
		semconv.DBSystemPostgreSQL,
		attribute.String("db.name", m.dbName),
	}
}

// RecordConnectionPoolStats records connection pool statistics
func (m *Metrics) RecordConnectionPoolStats(stats sql.DBStats) {
	if m.poolActiveConnections == nil {
		return
	}

	ctx := context.Background()
	attrs := m.commonAttributes()

	m.poolActiveConnections.Add(ctx, int64(stats.InUse), metric.WithAttributes(attrs...))
	m.poolIdleConnections.Add(ctx, int64(stats.Idle), metric.WithAttributes(attrs...))
	m.poolOpenConnections.Add(ctx, int64(stats.OpenConnections), metric.WithAttributes(attrs...))
	m.poolMaxOpenConnections.Add(ctx, int64(stats.MaxOpenConnections), metric.WithAttributes(attrs...))
	m.poolWaitCount.Add(ctx, stats.WaitCount, metric.WithAttributes(attrs...))
	m.poolWaitDuration.Record(ctx, float64(stats.WaitDuration.Nanoseconds())/1e9, metric.WithAttributes(attrs...))
	m.poolIdleClosed.Add(ctx, stats.MaxIdleClosed, metric.WithAttributes(attrs...))
	m.poolLifetimeClosed.Add(ctx, stats.MaxLifetimeClosed, metric.WithAttributes(attrs...))
}

// RecordQuery records query metrics
func (m *Metrics) RecordQuery(duration time.Duration, operation string, success bool) {
	if m.queryDuration == nil {
		return
	}

	ctx := context.Background()

	attrs := append(m.commonAttributes(), attribute.String("db.operation", operation))

	m.queryDuration.Record(ctx, duration.Seconds(), metric.WithAttributes(attrs...))
	m.queryCounter.Add(ctx, 1, metric.WithAttributes(attrs...))

	if !success {
		m.queryErrors.Add(ctx, 1, metric.WithAttributes(attrs...))
	}
}

// RecordHealthCheck records health check metrics
func (m *Metrics) RecordHealthCheck(duration time.Duration, success bool) {
	if m.healthCheckDuration == nil {
		return
	}

	ctx := context.Background()
	attrs := m.commonAttributes()

	m.healthCheckDuration.Record(ctx, duration.Seconds(), metric.WithAttributes(attrs...))
	m.healthCheckCounter.Add(ctx, 1, metric.WithAttributes(attrs...))

	if !success {
		m.healthCheckErrors.Add(ctx, 1, metric.WithAttributes(attrs...))
	}
}

// StartPeriodicStatsCollection starts periodic collection of connection pool stats
func (m *Metrics) StartPeriodicStatsCollection(db *sql.DB, interval time.Duration) func() {
	if m.poolActiveConnections == nil {
		return func() {} // No-op if pool metrics are disabled
	}

	ctx, cancel := context.WithCancel(context.Background())

	xsync.Go("metrics.postgresql.startPeriodicStatsCollection", func() error {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-ticker.C:
				stats := db.Stats()
				m.RecordConnectionPoolStats(stats)
			}
		}

		return nil
	})

	return cancel
}
