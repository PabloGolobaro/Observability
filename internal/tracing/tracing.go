package tracing

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

const Name = "Currency-api"

// newExporter returns a console exporter.
func newExporter() (trace.SpanExporter, error) {
	exporter, err := jaeger.New(
		jaeger.WithAgentEndpoint(jaeger.WithAgentHost("jaeger")))
	if err != nil {
		return nil, err
	}
	return exporter, nil
}
func newResource() *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("Currency"),
			semconv.ServiceVersionKey.String("v0.1.0"),
			attribute.String("environment", "demo"),
		),
	)
	return r
}

func TracerProvider() (*trace.TracerProvider, error) {
	exporter, err := newExporter()
	if err != nil {
		return nil, err
	}
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(newResource()),
	)
	return tp, nil
}
