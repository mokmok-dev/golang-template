//go:generate go run github.com/golang/mock/mockgen -source $GOFILE -package=$GOPACKAGE -destination=mock_$GOFILE

package tracer

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

type Tracer interface {
	Provider() trace.TracerProvider
	Tracer() trace.Tracer
	Shutdown(context.Context) error
}
