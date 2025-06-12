package postgresql

import (
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func TestNewMetrics(t *testing.T) {
	t.Parallel()

	config := DefaultMetricsConfig("postgres-metrics", "testdb")

	metrics, err := NewMetrics(config)
	require.NoError(t, err)
	require.NotNil(t, metrics)

	assert.Equal(t, "testdb", metrics.dbName)
	assert.Equal(t, "postgresql", metrics.dbSystem)
}

func TestMetricsConfig(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		config      *MetricsConfig
		expectError bool
	}{
		{
			name:        "nil config",
			config:      nil,
			expectError: true,
		},
		{
			name:        "valid config",
			config:      DefaultMetricsConfig("postgres-metrics", "testdb"),
			expectError: false,
		},
		{
			name: "custom config",
			config: &MetricsConfig{
				ServiceName:         "custom-service",
				DBName:              "customdb",
				DBSystem:            "postgresql",
				MeterProvider:       otel.GetMeterProvider(),
				StatsReportInterval: 10 * time.Second,
				EnableQueryMetrics:  true,
				EnablePoolMetrics:   false,
				EnableHealthMetrics: true,
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			metrics, err := NewMetrics(tt.config)

			if tt.expectError {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, metrics)
		})
	}
}

func TestMetricsRecordQuery(t *testing.T) {
	t.Parallel()

	config := DefaultMetricsConfig("postgres-metrics", "testdb")
	metrics, err := NewMetrics(config)
	require.NoError(t, err)

	// Test recording query metrics
	duration := 100 * time.Millisecond
	operation := "SELECT"

	// This should not panic
	metrics.RecordQuery(duration, operation, true)
	metrics.RecordQuery(duration, operation, false)
}

func TestMetricsRecordHealthCheck(t *testing.T) {
	t.Parallel()

	config := DefaultMetricsConfig("postgres-metrics", "testdb")
	metrics, err := NewMetrics(config)
	require.NoError(t, err)

	// Test recording health check metrics
	duration := 50 * time.Millisecond

	// This should not panic
	metrics.RecordHealthCheck(duration, true)
	metrics.RecordHealthCheck(duration, false)
}

func TestMetricsCommonAttributes(t *testing.T) {
	t.Parallel()

	config := DefaultMetricsConfig("postgres-metrics", "testdb")
	metrics, err := NewMetrics(config)
	require.NoError(t, err)

	attrs := metrics.commonAttributes()

	assert.Equal(t, 2, len(attrs))

	// Check for required attributes
	var foundDBSystem, foundDBName bool

	for _, attr := range attrs {
		switch attr.Key {
		case semconv.DBSystemKey:
			foundDBSystem = true

			assert.Equal(t, "postgresql", attr.Value.AsString())
		case "db.name":
			foundDBName = true

			assert.Equal(t, "testdb", attr.Value.AsString())
		}
	}

	assert.True(t, foundDBSystem, "Missing db.system attribute")
	assert.True(t, foundDBName, "Missing db.name attribute")
}

func TestDefaultMetricsConfig(t *testing.T) {
	t.Parallel()

	dbname := "testdb"
	config := DefaultMetricsConfig("postgres-metrics", dbname)

	require.NotNil(t, config)

	assert.Equal(t, dbname, config.DBName)
	assert.Equal(t, "postgresql", config.DBSystem)
	assert.Equal(t, "postgres-metrics", config.ServiceName)
	assert.True(t, config.EnableQueryMetrics)
	assert.True(t, config.EnablePoolMetrics)
	assert.True(t, config.EnableHealthMetrics)
	assert.Equal(t, 30*time.Second, config.StatsReportInterval)
}

func TestMetricsWithDisabledFeatures(t *testing.T) {
	t.Parallel()

	config := &MetricsConfig{
		ServiceName:         "test-service",
		DBName:              "testdb",
		DBSystem:            "postgresql",
		MeterProvider:       otel.GetMeterProvider(),
		StatsReportInterval: 10 * time.Second,
		EnableQueryMetrics:  false,
		EnablePoolMetrics:   false,
		EnableHealthMetrics: false,
	}

	metrics, err := NewMetrics(config)
	if err != nil {
		t.Fatalf("Failed to create metrics: %v", err)
	}

	// These should not panic even with disabled features
	metrics.RecordQuery(100*time.Millisecond, "SELECT", true)
	metrics.RecordHealthCheck(50*time.Millisecond, true)

	// Mock database stats
	mockStats := sql.DBStats{
		MaxOpenConnections: 10,
		OpenConnections:    5,
		InUse:              2,
		Idle:               3,
		WaitCount:          1,
		WaitDuration:       100 * time.Millisecond,
		MaxIdleClosed:      0,
		MaxLifetimeClosed:  0,
	}

	metrics.RecordConnectionPoolStats(mockStats)
}
