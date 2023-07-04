package usecase

import (
	"context"

	"github.com/mokmok-dev/golang-template/domain/model"
)

type GetUserByID interface {
	Do(context.Context, GetUserByIDInput) (*GetUserByIDOutput, error)
}

type GetUserByIDInput struct {
	ID string
}

type GetUserByIDOutput struct {
	User *model.User
}
