package postgresql

import (
	"database/sql"
	"fmt"

	"github.com/XSAM/otelsql"
	"github.com/go-pantheon/fabrica-util/data/db/postgresql"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

// PostgreSQLConfig holds comprehensive configuration for PostgreSQL connection with tracing
type PostgreSQLConfig struct {
	postgresql.Config

	IncludeQueryParameters bool
	DisableErrSkip         bool
	DisableQuery           bool
	CustomAttributes       []attribute.KeyValue
}

// DefaultPostgreSQLConfig returns a default configuration with sensible defaults
func DefaultPostgreSQLConfig(dbConfig postgresql.Config) *PostgreSQLConfig {
	return &PostgreSQLConfig{
		Config:                 dbConfig,
		IncludeQueryParameters: false,
		DisableErrSkip:         false,
		DisableQuery:           false,
		CustomAttributes: []attribute.KeyValue{
			semconv.DBSystemPostgreSQL,
			attribute.String("db.name", dbConfig.DBName),
		},
	}
}

func NewTracingDB(config *PostgreSQLConfig) (db *sql.DB, cleanup func(), err error) {
	var options []otelsql.Option

	if len(config.CustomAttributes) > 0 {
		options = append(options, otelsql.WithAttributes(config.CustomAttributes...))
	}

	if config.IncludeQueryParameters {
		options = append(options, otelsql.WithSQLCommenter(true))
	}

	// Configure span options
	spanOpts := otelsql.SpanOptions{
		DisableErrSkip: config.DisableErrSkip,
		DisableQuery:   config.DisableQuery,
	}
	options = append(options, otelsql.WithSpanOptions(spanOpts))

	driverName, err := otelsql.Register("postgres", options...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to register otel driver: %w", err)
	}

	return postgresql.New(driverName, config.Config)
}
