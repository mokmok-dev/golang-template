package configuration

import (
	"fmt"
	"strings"

	"github.com/google/wire"
	domain "github.com/mokmok-dev/golang-template/domain/configuration"
	"github.com/spf13/viper"
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
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to load environment variables: %w", err)
	}

	config := new(domain.Config)
	if err := viper.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal environment variable: %w", err)
	}

	return config, nil
}
