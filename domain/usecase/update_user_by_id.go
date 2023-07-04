package usecase

import (
	"context"

	"github.com/mokmok-dev/golang-template/domain/model"
)

type UpdateUserByID interface {
	Do(context.Context, UpdateUserByIDInput) (*UpdateUserByIDOutput, error)
}

type UpdateUserByIDInput struct {
	ID string
}

type UpdateUserByIDOutput struct {
	User *model.User
}
