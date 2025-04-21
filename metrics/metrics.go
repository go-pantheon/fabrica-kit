// Package metrics provides functionality for metrics collection and monitoring
// using OpenTelemetry for both server and client operations.
package metrics

import (
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

var (
	_metricRequests metric.Int64Counter
	_metricSeconds  metric.Float64Histogram
)

// Init initializes metrics collectors with the given service name.
// It sets up request counters and duration histograms.
func Init(name string) {
	meter := otel.Meter(name)

	var err error
	_metricRequests, err = metrics.DefaultRequestsCounter(meter, metrics.DefaultServerRequestsCounterName)

	if err != nil {
		panic(err)
	}

	_metricSeconds, err = metrics.DefaultSecondsHistogram(meter, metrics.DefaultServerSecondsHistogramName)

	if err != nil {
		panic(err)
	}
}

// Server returns a middleware that collects metrics for server operations.
func Server() middleware.Middleware {
	return metrics.Server(
		metrics.WithSeconds(_metricSeconds),
		metrics.WithRequests(_metricRequests),
	)
}

// Client returns a middleware that collects metrics for client operations.
func Client() middleware.Middleware {
	return metrics.Client(
		metrics.WithSeconds(_metricSeconds),
		metrics.WithRequests(_metricRequests),
	)
}
