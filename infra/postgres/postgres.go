package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/url"

	"github.com/google/wire"
	_ "github.com/lib/pq"
	"github.com/mokmok-dev/golang-template/domain/configuration"
	"github.com/mokmok-dev/golang-template/domain/logger"
)

var NewPostgresSet = wire.NewSet(
	NewPostgres,
)

func NewPostgres(
	ctx context.Context,
	logger logger.Logger,
	config configuration.Database,
) (*sql.DB, error) {
	dsn := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(config.User, config.Password),
		Host:   net.JoinHostPort(config.Host, config.Port),
		Path:   config.Name,
		RawQuery: url.Values{
			"sslmode": []string{"disable"},
		}.Encode(),
	}

	logger.Debug(ctx, "start postgres connector")

	db, err := sql.Open("postgres", dsn.String())
	if err != nil {
		return nil, fmt.Errorf("failed to make connection for postgres: %w", err)
	}

	return db, nil
}
