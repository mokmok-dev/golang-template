package handler

import (
	"context"
	"fmt"

	"github.com/bufbuild/connect-go"
	"github.com/google/wire"
	"github.com/mokmok-dev/golang-template/adapter/handler/marshaller"
	"github.com/mokmok-dev/golang-template/domain/logger"
	"github.com/mokmok-dev/golang-template/domain/tracer"
	"github.com/mokmok-dev/golang-template/domain/usecase"
	v1 "github.com/mokmok-dev/golang-template/proto/golang-template/v1"
	"github.com/mokmok-dev/golang-template/proto/golang-template/v1/v1connect"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ v1connect.UserServiceHandler = (*User)(nil)

var NewUserSet = wire.NewSet(
	wire.Bind(new(v1connect.UserServiceHandler), new(*User)),
	NewUser,
)

type User struct {
	logger         logger.Logger
	tracer         tracer.Tracer
	createUser     usecase.CreateUser
	getUserByID    usecase.GetUserByID
	updateUserByID usecase.UpdateUserByID
	removeUserByID usecase.RemoveUserByID
}

func NewUser(
	logger logger.Logger,
	tracer tracer.Tracer,
	createUser usecase.CreateUser,
	getUserByID usecase.GetUserByID,
	updateUserByID usecase.UpdateUserByID,
	removeUserByID usecase.RemoveUserByID,
) *User {
	return &User{
		logger:         logger,
		tracer:         tracer,
		createUser:     createUser,
		getUserByID:    getUserByID,
		updateUserByID: updateUserByID,
		removeUserByID: removeUserByID,
	}
}
func (u *User) CreateUser(ctx context.Context, req *connect.Request[v1.CreateUserRequest]) (*connect.Response[v1.CreateUserResponse], error) {
	ctx, span := u.tracer.Tracer().Start(ctx, "handler.CreateUser")
	defer span.End()

	result, err := u.createUser.Do(ctx, usecase.CreateUserInput{})
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &connect.Response[v1.CreateUserResponse]{
		Msg: &v1.CreateUserResponse{
			User: marshaller.UserToProto(result.User),
		},
	}, nil
}

func (u *User) GetUserByID(ctx context.Context, req *connect.Request[v1.GetUserByIDRequest]) (*connect.Response[v1.GetUserByIDResponse], error) {
	ctx, span := u.tracer.Tracer().Start(ctx, "handler.GetUserByID")
	defer span.End()

	result, err := u.getUserByID.Do(ctx, usecase.GetUserByIDInput{
		ID: req.Msg.GetId(),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}

	return &connect.Response[v1.GetUserByIDResponse]{
		Msg: &v1.GetUserByIDResponse{
			User: marshaller.UserToProto(result.User),
		},
	}, nil
}

func (u *User) UpdateUserByID(ctx context.Context, req *connect.Request[v1.UpdateUserByIDRequest]) (*connect.Response[v1.UpdateUserByIDResponse], error) {
	ctx, span := u.tracer.Tracer().Start(ctx, "handler.UpdateUserByID")
	defer span.End()

	result, err := u.updateUserByID.Do(ctx, usecase.UpdateUserByIDInput{
		ID: req.Msg.GetId(),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update user by ID: %w", err)
	}

	return &connect.Response[v1.UpdateUserByIDResponse]{
		Msg: &v1.UpdateUserByIDResponse{
			User: marshaller.UserToProto(result.User),
		},
	}, nil
}

func (u *User) RemoveUserByID(ctx context.Context, req *connect.Request[v1.RemoveUserByIDRequest]) (*connect.Response[emptypb.Empty], error) {
	ctx, span := u.tracer.Tracer().Start(ctx, "handler.RemoveUserByID")
	defer span.End()

	if err := u.removeUserByID.Do(ctx, usecase.RemoveUserByIDInput{
		ID: req.Msg.GetId(),
	}); err != nil {
		return nil, fmt.Errorf("failed to remove user by ID: %w", err)
	}

	return &connect.Response[emptypb.Empty]{
		Msg: &emptypb.Empty{},
	}, nil
}
