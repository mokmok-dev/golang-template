package repository

import (
	context "context"

	"github.com/mokmok-dev/golang-template/domain/model"
)

type User interface {
	CreateUser(context.Context, *model.User) (*model.User, error)
	GetUsersByID(context.Context, model.UserID) (*model.User, error)
	UpdateUserByID(context.Context, *model.User) (*model.User, error)
	RemoveUserByID(context.Context, model.UserID) error
}
