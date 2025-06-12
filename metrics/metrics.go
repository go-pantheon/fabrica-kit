// Package metrics provides functionality for metrics collection and monitoring
// using OpenTelemetry for both server and client operations.
package metrics

import (
	"sync/atomic"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

var (
	_metricRequests atomic.Value // metric.Int64Counter
	_metricSeconds  atomic.Value // metric.Float64Histogram
)

func init() {
	Init("")
}

// Init initializes metrics collectors with the given service name.
// It sets up request counters and duration histograms.
func Init(name string) {
	meter := otel.Meter(name)

	req, err := metrics.DefaultRequestsCounter(meter, metrics.DefaultServerRequestsCounterName)
	if err != nil {
		panic(err)
	}

	_metricRequests.Store(req)

	sec, err := metrics.DefaultSecondsHistogram(meter, metrics.DefaultServerSecondsHistogramName)
	if err != nil {
		panic(err)
	}

	_metricSeconds.Store(sec)
}

// Server returns a middleware that collects metrics for server operations.
func Server() middleware.Middleware {
	return metrics.Server(
		metrics.WithSeconds(_metricSeconds.Load().(metric.Float64Histogram)),
		metrics.WithRequests(_metricRequests.Load().(metric.Int64Counter)),
	)
}

// Client returns a middleware that collects metrics for client operations.
func Client() middleware.Middleware {
	return metrics.Client(
		metrics.WithSeconds(_metricSeconds.Load().(metric.Float64Histogram)),
		metrics.WithRequests(_metricRequests.Load().(metric.Int64Counter)),
	)
}
