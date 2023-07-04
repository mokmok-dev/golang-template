package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/wire"
	"github.com/mokmok-dev/golang-template/domain/logger"
	"github.com/mokmok-dev/golang-template/domain/model"
	"github.com/mokmok-dev/golang-template/domain/repository"
	"github.com/mokmok-dev/golang-template/domain/tracer"
	domain "github.com/mokmok-dev/golang-template/domain/usecase"
)

var _ domain.UpdateUserByID = (*UpdateUserByID)(nil)

var NewUpdateUserByIDSet = wire.NewSet(
	wire.Bind(new(domain.UpdateUserByID), new(*UpdateUserByID)),
	NewUpdateUserByID,
)

type UpdateUserByID struct {
	logger     logger.Logger
	tracer     tracer.Tracer
	repository repository.Repository
}

func NewUpdateUserByID(
	logger logger.Logger,
	tracer tracer.Tracer,
	repository repository.Repository,
) *UpdateUserByID {
	return &UpdateUserByID{
		logger:     logger,
		tracer:     tracer,
		repository: repository,
	}
}

func (u *UpdateUserByID) Do(ctx context.Context, input domain.UpdateUserByIDInput) (*domain.UpdateUserByIDOutput, error) {
	ctx, span := u.tracer.Tracer().Start(ctx, "usecase.UpdateUserByID")
	defer span.End()

	id, err := model.ParseUserID(input.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse user ID: %w", err)
	}

	origin, err := u.repository.GetUsersByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("faield to get user: %w", err)
	}

	updated := u.updateUser(origin, input)

	user, err := u.repository.UpdateUserByID(ctx, updated)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return &domain.UpdateUserByIDOutput{
		User: user,
	}, nil
}

func (u *UpdateUserByID) updateUser(user *model.User, input domain.UpdateUserByIDInput) *model.User {
	user.UpdatedAt = time.Now().UTC()

	return user
}
