package model

import "github.com/google/wire"

var NewConnSet = wire.NewSet(
	wire.Bind(new(Querier), new(*Queries)),
	New,
)
