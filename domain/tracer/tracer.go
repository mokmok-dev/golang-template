package tracer

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

type Tracer interface {
	Tracer() trace.Tracer
	Shutdown(context.Context) error
}
