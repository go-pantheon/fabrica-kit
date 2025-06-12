package trace

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

// Init initializes the OpenTelemetry trace provider with the specified configuration.
// It sets up an exporter to the provided URL and configures the tracer with service name,
// profile, and color attributes.
func Init(url string, name string, profile, color string) error {
	exporter, err := otlptracehttp.New(context.Background(),
		otlptracehttp.WithEndpoint(url),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		return err
	}

	tp := tracesdk.NewTracerProvider(
		tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.TraceIDRatioBased(1.0))),
		tracesdk.WithBatcher(exporter),
		tracesdk.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(name),
			attribute.String("profile", profile),
			attribute.String("color", color),
		)),
	)
	otel.SetTracerProvider(tp)

	return nil
}
