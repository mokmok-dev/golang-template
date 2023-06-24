package configuration

import (
	"fmt"

	"github.com/google/wire"
	"github.com/kelseyhightower/envconfig"
	domain "github.com/mokmok-dev/golang-template/domain/configuration"
)

var NewConfigSet = wire.NewSet(
	wire.FieldsOf(new(*domain.Config), "Log"),
	wire.FieldsOf(new(*domain.Config), "Server"),
	wire.FieldsOf(new(*domain.Config), "Database"),
	wire.FieldsOf(new(*domain.Config), "Email"),
	wire.FieldsOf(new(*domain.Config), "GCP"),
	NewConfig,
)

func NewConfig() (*domain.Config, error) {
	config := new(domain.Config)

	if err := envconfig.Process("", config); err != nil {
		return nil, fmt.Errorf("failed to load environment variables: %w", err)
	}

	return config, nil
}
