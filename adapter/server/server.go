package server

import (
	"net/http"
	"time"

	"github.com/google/wire"
	"github.com/mokmok-dev/golang-template/domain/configuration"
	"github.com/mokmok-dev/golang-template/domain/logger"
	"github.com/mokmok-dev/golang-template/domain/tracer"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

var NewServerSet = wire.NewSet(
	NewServer,
)

func NewServer(
	logger logger.Logger,
	tracer tracer.Tracer,
	config configuration.Server,
	handler http.Handler,
) *http.Server {
	return &http.Server{
		Addr: ":" + config.Port,
		Handler: otelhttp.NewHandler(handler,
			"server",
			otelhttp.WithTracerProvider(
				tracer.Provider(),
			),
			otelhttp.WithMessageEvents(
				otelhttp.ReadEvents,
				otelhttp.WriteEvents,
			),
		),
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
	}
}
