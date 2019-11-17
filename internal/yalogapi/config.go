package yalogapi

import (
	"github.com/d2one/yalogapi/internal/clickhouse"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Config struct {
	Token      string                       `mapstructure:"token"`
	Clickhouse *clickhouse.ClickhouseConfig `mapstructure:"clickhouse"`
	// logsapi    *logsapi.Config
}

func NewYalogapiConfig() (*Config, error) {
	return &Config{}, nil
}

func (config Config) Validate() error {
	return validation.ValidateStruct(&config,
		// Name cannot be empty, and the length must be between 5 and 20.
		validation.Field(&config.Token, validation.Required),
		validation.Field(&config.Clickhouse),
	)
}
