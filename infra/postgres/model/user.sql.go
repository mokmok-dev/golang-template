// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: user.sql

package model

import (
	"context"
	"time"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at
`

type CreateUserParams struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) CreateUser(ctx context.Context, arg *CreateUserParams) (*User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.ID, arg.CreatedAt, arg.UpdatedAt)
	var i User
	err := row.Scan(&i.ID, &i.CreatedAt, &i.UpdatedAt)
	return &i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, created_at, updated_at FROM users WHERE id = $1
`

func (q *Queries) GetUserByID(ctx context.Context, id string) (*User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(&i.ID, &i.CreatedAt, &i.UpdatedAt)
	return &i, err
}

const removeUserByID = `-- name: RemoveUserByID :exec
DELETE FROM users WHERE id = $1
`

func (q *Queries) RemoveUserByID(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, removeUserByID, id)
	return err
}

const updateUserByID = `-- name: UpdateUserByID :one
UPDATE users SET updated_at = $1 WHERE id = $2 RETURNING id, created_at, updated_at
`

type UpdateUserByIDParams struct {
	UpdatedAt time.Time
	ID        string
}

func (q *Queries) UpdateUserByID(ctx context.Context, arg *UpdateUserByIDParams) (*User, error) {
	row := q.db.QueryRowContext(ctx, updateUserByID, arg.UpdatedAt, arg.ID)
	var i User
	err := row.Scan(&i.ID, &i.CreatedAt, &i.UpdatedAt)
	return &i, err
}
