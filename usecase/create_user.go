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

var _ domain.CreateUser = (*CreateUser)(nil)

var NewCreateUserSet = wire.NewSet(
	wire.Bind(new(domain.CreateUser), new(*CreateUser)),
	NewCreateUser,
)

type CreateUser struct {
	logger     logger.Logger
	tracer     tracer.Tracer
	repository repository.Repository
}

func NewCreateUser(
	logger logger.Logger,
	tracer tracer.Tracer,
	repository repository.Repository,
) *CreateUser {
	return &CreateUser{
		logger:     logger,
		tracer:     tracer,
		repository: repository,
	}
}

func (u *CreateUser) Do(ctx context.Context, input domain.CreateUserInput) (*domain.CreateUserOutput, error) {
	ctx, span := u.tracer.Tracer().Start(ctx, "usecase.CreateUser")
	defer span.End()

	newUser := model.NewUser()

	user, err := u.repository.CreateUser(ctx, newUser)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &domain.CreateUserOutput{
		User: user,
	}, nil
}
