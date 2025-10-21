package trace

import (
	"context"
	"fmt"
	"log"

	"github.com/BaiMeow/NetworkMonitor/conf"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.34.0"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
)

var Tracer trace.Tracer = &noop.Tracer{}

func Init() error {
	exporter, err := newExporter()
	if err != nil {
		return fmt.Errorf("init Tracer fail:%v", err)
	}
	tp := newTraceProvider(exporter)
	otel.SetTracerProvider(tp)
	Tracer = tp.Tracer("DN11NetworkMonitor")
	return nil
}

func newExporter() (sdktrace.SpanExporter, error) {
	if conf.Trace.Endpoint == "" {
		log.Println("WARNING: no trace endpoint defined")
		return stdouttrace.New()
	}

	return otlptracegrpc.New(
		context.Background(),
		otlptracegrpc.WithEndpoint(conf.Trace.Endpoint),
		otlptracegrpc.WithInsecure(),
	)
}

func newTraceProvider(exp sdktrace.SpanExporter) *sdktrace.TracerProvider {
	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(conf.Trace.ServiceName),
		),
	)

	if err != nil {
		panic(err)
	}

	return sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(r),
	)
}
