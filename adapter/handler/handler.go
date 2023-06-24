package handler

import (
	"net/http"

	"github.com/google/wire"
	"github.com/mokmok-dev/golang-template/adapter/handler/middleware"
	"github.com/mokmok-dev/golang-template/domain/logger"
	"github.com/mokmok-dev/golang-template/domain/tracer"
)

var _ http.Handler = (*Handler)(nil)

var NewHandlerSet = wire.NewSet(
	wire.Bind(new(http.Handler), new(*Handler)),
	NewHandler,
)

type Handler struct {
	logger logger.Logger
	tracer tracer.Tracer
	router http.Handler
}

func NewHandler(
	logger logger.Logger,
	tracer tracer.Tracer,
) *Handler {
	mux := http.NewServeMux()

	router := middleware.Chain(mux,
		middleware.Logger(logger),
	)

	h := &Handler{
		logger: logger,
		tracer: tracer,
		router: router,
	}

	mux.HandleFunc("/healthcheck", h.healthcheck)

	return h
}

func (h *Handler) healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}
