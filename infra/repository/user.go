package repository

import (
	"context"
	"fmt"

	domainmodel "github.com/mokmok-dev/golang-template/domain/model"
	"github.com/mokmok-dev/golang-template/infra/postgres/marshaller"
	recordmodel "github.com/mokmok-dev/golang-template/infra/postgres/model"
)

func (r *Repository) CreateUser(ctx context.Context, model *domainmodel.User) (*domainmodel.User, error) {
	ctx, span := r.tracer.Tracer().Start(ctx, "repository.CreateUser")
	defer span.End()

	record, err := r.queries.CreateUser(ctx, &recordmodel.CreateUserParams{
		ID:        model.ID.String(),
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return marshaller.RecordToUser(record)
}

func (r *Repository) GetUsersByID(ctx context.Context, id domainmodel.UserID) (*domainmodel.User, error) {
	ctx, span := r.tracer.Tracer().Start(ctx, "repository.GetUsersByID")
	defer span.End()

	record, err := r.queries.GetUserByID(ctx, id.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}

	return marshaller.RecordToUser(record)
}

func (r *Repository) UpdateUserByID(ctx context.Context, model *domainmodel.User) (*domainmodel.User, error) {
	ctx, span := r.tracer.Tracer().Start(ctx, "repository.UpdateUserByID")
	defer span.End()

	record, err := r.queries.UpdateUserByID(ctx, &recordmodel.UpdateUserByIDParams{
		UpdatedAt: model.UpdatedAt,
		ID:        model.ID.String(),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update user by ID: %w", err)
	}

	return marshaller.RecordToUser(record)
}

func (r *Repository) RemoveUserByID(ctx context.Context, id domainmodel.UserID) error {
	ctx, span := r.tracer.Tracer().Start(ctx, "repository.RemoveUserByID")
	defer span.End()

	if err := r.queries.RemoveUserByID(ctx, id.String()); err != nil {
		return fmt.Errorf("failed to remove user by ID: %w", err)
	}

	return nil
}
