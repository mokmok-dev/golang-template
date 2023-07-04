package usecase

import (
	"context"
	"fmt"

	"github.com/google/wire"
	"github.com/mokmok-dev/golang-template/domain/logger"
	"github.com/mokmok-dev/golang-template/domain/model"
	"github.com/mokmok-dev/golang-template/domain/repository"
	"github.com/mokmok-dev/golang-template/domain/tracer"
	domain "github.com/mokmok-dev/golang-template/domain/usecase"
)

var _ domain.GetUserByID = (*GetUserByID)(nil)

var NewGetUserByIDSet = wire.NewSet(
	wire.Bind(new(domain.GetUserByID), new(*GetUserByID)),
	NewGetUserByID,
)

type GetUserByID struct {
	logger     logger.Logger
	tracer     tracer.Tracer
	repository repository.Repository
}

func NewGetUserByID(
	logger logger.Logger,
	tracer tracer.Tracer,
	repository repository.Repository,
) *GetUserByID {
	return &GetUserByID{
		logger:     logger,
		tracer:     tracer,
		repository: repository,
	}
}

func (u *GetUserByID) Do(ctx context.Context, input domain.GetUserByIDInput) (*domain.GetUserByIDOutput, error) {
	ctx, span := u.tracer.Tracer().Start(ctx, "usecase.GetUserByID")
	defer span.End()

	id, err := model.ParseUserID(input.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse user ID: %w", err)
	}

	user, err := u.repository.GetUsersByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &domain.GetUserByIDOutput{
		User: user,
	}, nil
}
