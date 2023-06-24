package tracer

import (
	"context"
	"fmt"

	exporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	gcppropagator "github.com/GoogleCloudPlatform/opentelemetry-operations-go/propagator"
	"github.com/google/wire"
	"github.com/mokmok-dev/golang-template/domain/configuration"
	domain "github.com/mokmok-dev/golang-template/domain/tracer"
	"go.opentelemetry.io/contrib/detectors/gcp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.9.0"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/multierr"
)

var _ domain.Tracer = (*Tracer)(nil)

var NewTracerSet = wire.NewSet(
	wire.Bind(new(domain.Tracer), new(*Tracer)),
	NewTracer,
)

type Tracer struct {
	exporter *exporter.Exporter
	provider *sdktrace.TracerProvider
}

func NewTracer(
	config configuration.GCP,
) (*Tracer, error) {
	ex, err := exporter.New(exporter.WithProjectID(config.ProjectID))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize exporter: %w", err)
	}

	res, err := resource.New(context.Background(),
		resource.WithDetectors(gcp.NewDetector()),
		resource.WithTelemetrySDK(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String("golang-template"),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize resource: %w", err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(ex),
		sdktrace.WithResource(res),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			gcppropagator.CloudTraceOneWayPropagator{},
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	return &Tracer{
		exporter: ex,
		provider: tp,
	}, nil
}

func (t *Tracer) Provider() trace.TracerProvider {
	return t.provider
}

func (t *Tracer) Tracer() trace.Tracer {
	return t.provider.Tracer("")
}

func (t *Tracer) Shutdown(ctx context.Context) error {
	var merr error

	defer multierr.AppendInto(&merr, t.provider.ForceFlush(ctx))
	defer multierr.AppendInto(&merr, t.provider.Shutdown(ctx))

	return merr
}
