package model

import (
	"time"
)

type User struct {
	ID        UserID
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser() *User {
	now := time.Now().UTC()

	return &User{
		ID:        NewUserID(),
		CreatedAt: now,
		UpdatedAt: now,
	}
}
