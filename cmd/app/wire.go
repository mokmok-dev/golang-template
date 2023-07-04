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
	"github.com/mokmok-dev/golang-template/infra/postgres"
	"github.com/mokmok-dev/golang-template/infra/postgres/model"
	"github.com/mokmok-dev/golang-template/infra/repository"
	"github.com/mokmok-dev/golang-template/infra/tracer"
	"github.com/mokmok-dev/golang-template/usecase"
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
		handler.NewUserSet,
		server.NewServerSet,
		usecase.NewCreateUserSet,
		usecase.NewGetUserByIDSet,
		usecase.NewUpdateUserByIDSet,
		usecase.NewRemoveUserByIDSet,
		repository.NewRepositorySet,
		model.NewConnSet,
		postgres.NewPostgresSet,

		wire.Struct(new(app), "*"),
	)

	return nil, nil
}
