//go:generate go run github.com/golang/mock/mockgen -source $GOFILE -package=$GOPACKAGE -destination=mock_$GOFILE

package repository

import "context"

type Repository interface {
	WithTx(context.Context, func(context.Context, Repository) error) error
}
