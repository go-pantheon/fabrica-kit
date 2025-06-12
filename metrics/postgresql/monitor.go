package postgresql

import (
	"context"
	"database/sql"
	"sync"
	"time"
)

// Monitor provides a convenient wrapper for monitoring database operations
type Monitor struct {
	metrics *Metrics
	db      *sql.DB
	mu      sync.RWMutex
	stopFn  func()
}

// WithMonitor creates a new database monitor
func WithMonitor(db *sql.DB, config *MetricsConfig) (*Monitor, error) {
	metrics, err := NewMetrics(config)
	if err != nil {
		return nil, err
	}

	monitor := &Monitor{
		metrics: metrics,
		db:      db,
	}

	// Start periodic stats collection if pool metrics are enabled
	if config.EnablePoolMetrics {
		stopFn := metrics.StartPeriodicStatsCollection(db, config.StatsReportInterval)
		monitor.stopFn = stopFn
	}

	return monitor, nil
}

// QueryContext wraps sql.DB.QueryContext with metrics
func (dm *Monitor) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	start := time.Now()

	rows, err := dm.db.QueryContext(ctx, query, args...)

	duration := time.Since(start)
	success := err == nil
	dm.metrics.RecordQuery(duration, "query", success)

	return rows, err
}

// ExecContext wraps sql.DB.ExecContext with metrics
func (dm *Monitor) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	start := time.Now()

	result, err := dm.db.ExecContext(ctx, query, args...)

	duration := time.Since(start)
	success := err == nil
	dm.metrics.RecordQuery(duration, "exec", success)

	return result, err
}

// HealthCheck performs a health check with metrics
func (dm *Monitor) HealthCheck(ctx context.Context) error {
	start := time.Now()

	err := dm.db.PingContext(ctx)

	duration := time.Since(start)
	success := err == nil
	dm.metrics.RecordHealthCheck(duration, success)

	return err
}

// Close stops the monitor and releases resources
func (dm *Monitor) Close() {
	dm.mu.Lock()
	defer dm.mu.Unlock()

	if dm.stopFn != nil {
		dm.stopFn()
		dm.stopFn = nil
	}
}

// DB returns the underlying database connection
func (dm *Monitor) DB() *sql.DB {
	return dm.db
}
