package main

import (
	"os"
	"os/signal"
	"time"

	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"

	"github.com/Coflnet/sky-controller/internal/metrics"
	"github.com/Coflnet/sky-controller/internal/usecase"
)

func main() {

  tp, err := tracerProvider()
	if err != nil {
		log.Panic().Err(err).Msg("failed to create tracer provider")
	}

	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(tp)


  go metrics.Start()

	proxyScaler := usecase.ProxyScaler{
		Interval: 1 * time.Minute,
	}
  proxyScaler.Start()

  activeSubscriptionsWatcher := usecase.ActiveSubscriptionsWatcher{
    Interval: 1 * time.Minute,
    ProductUpdateInterval: 1 * time.Hour,
  }
  activeSubscriptionsWatcher.Start()

	// wait for exit signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Info().Msg("Shutting down...")
}

func newResource() *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("sky-controller"),
			semconv.ServiceVersionKey.String("v0.0.1"),
		),
	)
	return r
}

func tracerProvider() (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithAgentEndpoint())
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
      semconv.ServiceNameKey.String("sky-controller"),
      semconv.ServiceVersionKey.String("v0.0.1"),
		)),
	)
	return tp, nil
}
