package postgresql

import (
	"context"

	"github.com/exaring/otelpgx"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-pantheon/fabrica-util/data/db/postgresql"
	"github.com/go-pantheon/fabrica-util/errors"
	"github.com/jackc/pgx/v5"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

// PostgreSQLConfig holds comprehensive configuration for PostgreSQL connection pool with tracing
type PostgreSQLConfig struct {
	postgresql.Config

	// OpenTelemetry tracing options
	IncludeQueryParameters bool
	DisableSQLInAttributes bool // Use WithDisableSQLStatementInAttributes
	DisableMetrics         bool // Control metrics collection separately
	CustomAttributes       []attribute.KeyValue
}

// DefaultPostgreSQLConfig returns a default configuration with sensible defaults
func DefaultPostgreSQLConfig(dbConfig postgresql.Config) *PostgreSQLConfig {
	return &PostgreSQLConfig{
		Config:                 dbConfig,
		IncludeQueryParameters: false,
		DisableSQLInAttributes: false,
		DisableMetrics:         false,
		CustomAttributes: []attribute.KeyValue{
			semconv.DBSystemPostgreSQL,
			attribute.String("db.name", dbConfig.DBName),
		},
	}
}

func NewTracingDB(ctx context.Context, config *PostgreSQLConfig) (db *postgresql.DB, cleanup func(), err error) {
	// Create tracer with options
	var opts []otelpgx.Option

	if len(config.CustomAttributes) > 0 {
		opts = append(opts, otelpgx.WithTracerAttributes(config.CustomAttributes...))
	}

	if config.IncludeQueryParameters {
		opts = append(opts, otelpgx.WithIncludeQueryParameters())
	}

	if config.DisableSQLInAttributes {
		opts = append(opts, otelpgx.WithDisableSQLStatementInAttributes())
	}

	poolConfig := postgresql.NewConfig(config.DSN, config.DBName)
	poolConfig.Tracer = otelpgx.NewTracer(opts...)

	db, cleanup, err = postgresql.NewDBFromConfig(poolConfig)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to create connection pool")
	}

	if !config.DisableMetrics {
		if err := otelpgx.RecordStats(db.GetPool()); err != nil {
			log.Warnf("Warning: failed to record pool stats: %v\n", err)
		}
	}

	return db, cleanup, nil
}

// CreateTracer creates a pgx.QueryTracer for use with existing pools
func CreateTracer(opts ...otelpgx.Option) pgx.QueryTracer {
	return otelpgx.NewTracer(opts...)
}
