//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"net/http"

	"github.com/google/wire"
	"github.com/mokmok-dev/golang-template/adapter/handler"
	"github.com/mokmok-dev/golang-template/adapter/server"
	domainlogger "github.com/mokmok-dev/golang-template/domain/logger"
	"github.com/mokmok-dev/golang-template/infra/configuration"
	"github.com/mokmok-dev/golang-template/infra/logger"
	"github.com/mokmok-dev/golang-template/infra/tracer"
)

type app struct {
	ctx    context.Context
	logger domainlogger.Logger
	server *http.Server
}

func initialize() (*app, error) {
	wire.Build(
		context.Background,
		configuration.NewConfigSet,
		tracer.NewTracerSet,
		logger.NewLoggerSet,
		handler.NewHandlerSet,
		server.NewServerSet,

		wire.Struct(new(app), "*"),
	)

	return nil, nil
}
