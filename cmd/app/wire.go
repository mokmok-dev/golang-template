//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
)

type app struct{}

func initialize() (*app, error) {
	wire.Build(
		wire.Struct(new(app), "*"),
	)
	return nil, nil
}
