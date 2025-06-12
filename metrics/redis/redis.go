package redis

import (
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
)

// WithMetrics adds OpenTelemetry metrics to a Cacheable interface from fabrica-util
func WithMetrics(client redis.UniversalClient, config *MetricsConfig) error {
	if config == nil {
		config = DefaultMetricsConfig()
	}

	return redisotel.InstrumentMetrics(
		client,
		redisotel.WithMeterProvider(config.MeterProvider),
	)
}
