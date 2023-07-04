package usecase

import "context"

type RemoveUserByID interface {
	Do(context.Context, RemoveUserByIDInput) error
}

type RemoveUserByIDInput struct {
	ID string
}
