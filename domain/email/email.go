//go:generate go run github.com/golang/mock/mockgen -source $GOFILE -package=$GOPACKAGE -destination=mock_$GOFILE

package email

type Email interface{}
