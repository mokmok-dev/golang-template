package repository

import "context"

type Repository interface {
	WithTx(context.Context, func(context.Context, Repository) error) error
}
