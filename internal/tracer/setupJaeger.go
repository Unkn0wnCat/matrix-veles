package tracer

import (
	"context"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
	"log"
	"os"
	"runtime"
	"runtime/debug"
)

var (
	TraceProvider *tracesdk.TracerProvider
	Ctx           context.Context
	Cancel        context.CancelFunc
	Tracer        trace.Tracer
)

func SetupDummy() {
	tp := tracesdk.NewTracerProvider()
	TraceProvider = tp
	Ctx, Cancel = context.WithCancel(context.Background())
	Tracer = tp.Tracer("matrix_veles")
}

func SetupJaeger() {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(viper.GetString("tracing.jaeger.endpoint"))))
	if err != nil {
		log.Fatal(err)
	}

	bi, ok := debug.ReadBuildInfo()
	if !ok {
		log.Fatal(ok)
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("matrix_veles"),
			attribute.String("version", bi.Main.Version),
			attribute.String("go_version", runtime.Version()),
			attribute.String("hostname", hostname),
			attribute.String("os", runtime.GOOS),
			attribute.String("arch", runtime.GOARCH),
		)),
	)

	otel.SetTracerProvider(tp)

	TraceProvider = tp
	Ctx, Cancel = context.WithCancel(context.Background())
	Tracer = tp.Tracer("matrix_veles")
}
