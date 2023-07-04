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

var _ domain.RemoveUserByID = (*RemoveUserByID)(nil)

var NewRemoveUserByIDSet = wire.NewSet(
	wire.Bind(new(domain.RemoveUserByID), new(*RemoveUserByID)),
	NewRemoveUserByID,
)

type RemoveUserByID struct {
	logger     logger.Logger
	tracer     tracer.Tracer
	repository repository.Repository
}

func NewRemoveUserByID(
	logger logger.Logger,
	tracer tracer.Tracer,
	repository repository.Repository,
) *RemoveUserByID {
	return &RemoveUserByID{
		logger:     logger,
		tracer:     tracer,
		repository: repository,
	}
}

func (u *RemoveUserByID) Do(ctx context.Context, input domain.RemoveUserByIDInput) error {
	ctx, span := u.tracer.Tracer().Start(ctx, "usecase.RemoveUserByID")
	defer span.End()

	id, err := model.ParseUserID(input.ID)
	if err != nil {
		return fmt.Errorf("failed to parse user ID: %w", err)
	}

	if err := u.repository.RemoveUserByID(ctx, id); err != nil {
		return fmt.Errorf("failed to remove user: %w", err)
	}

	return nil
}
