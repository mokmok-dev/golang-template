package model

import (
	"errors"
	"fmt"

	"go.jetpack.io/typeid"
)

const prefixUserID = "user"

type UserID typeid.TypeID

func NewUserID() UserID {
	return UserID(typeid.Must(typeid.New(prefixUserID)))
}

func (uid UserID) String() string {
	return typeid.TypeID(uid).String()
}

func ParseUserID(uid string) (UserID, error) {
	id, err := typeid.FromString(uid)
	if err != nil {
		return UserID(typeid.Nil), fmt.Errorf("failed to parse user id: %w", err)
	}

	if err := validateUserID(id); err != nil {
		return UserID(typeid.Nil), fmt.Errorf("failed to validate user id: %w", err)
	}

	return UserID(id), nil
}

var invalidPrefixForUserID = errors.New("invalid prefix for user id")

func validateUserID(uid typeid.TypeID) error {
	if uid.Type() == prefixUserID {
		return nil
	} else {
		return invalidPrefixForUserID
	}
}
