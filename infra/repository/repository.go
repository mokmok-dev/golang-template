package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/wire"
	domain "github.com/mokmok-dev/golang-template/domain/repository"
	"github.com/mokmok-dev/golang-template/domain/tracer"
	"github.com/mokmok-dev/golang-template/infra/postgres/model"
)

var _ domain.Repository = (*Repository)(nil)

var NewRepositorySet = wire.NewSet(
	wire.Bind(new(domain.Repository), new(*Repository)),
	NewRepository,
)

type Repository struct {
	tracer  tracer.Tracer
	db      *sql.DB
	queries *model.Queries
}

func NewRepository(
	tracer tracer.Tracer,
	db *sql.DB,
	queries *model.Queries,
) *Repository {
	return &Repository{
		tracer:  tracer,
		db:      db,
		queries: queries,
	}
}

func (r *Repository) WithTx(ctx context.Context, inner func(context.Context, domain.Repository) error) error {
	ctx, span := r.tracer.Tracer().Start(ctx, "repository.WithTx")
	defer span.End()

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			if rerr := tx.Rollback(); rerr != nil {
				err = fmt.Errorf("failed to rollback: %w", rerr)
			}

			panic(p)
		}

		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				err = fmt.Errorf("failed to rollback: %w", rerr)
			}
		} else {
			if cerr := tx.Commit(); cerr != nil {
				err = fmt.Errorf("failed to commit: %w", cerr)
			}
		}
	}()

	repo := &Repository{
		tracer:  r.tracer,
		db:      r.db,
		queries: r.queries.WithTx(tx),
	}

	if err := inner(ctx, repo); err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}
