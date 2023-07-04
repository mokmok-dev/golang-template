package usecase

import (
	"context"

	"github.com/mokmok-dev/golang-template/domain/model"
)

type CreateUser interface {
	Do(context.Context, CreateUserInput) (*CreateUserOutput, error)
}

type CreateUserInput struct {
}

type CreateUserOutput struct {
	User *model.User
}
