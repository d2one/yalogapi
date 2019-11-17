package yalogapi

import (
	"github.com/d2one/yalogapi/internal/clickhouse"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Config struct {
	StartDate  string
	EndDate    string
	Mode       string
	Source     string
	Clickhouse *clickhouse.ClickhouseConfig `mapstructure:"clickhouse"`
	// logsapi    *logsapi.Config
}

func NewYalogapiConfig() (*Config, error) {
	return &Config{}, nil
}

func (config Config) Validate() error {
	return validation.ValidateStruct(&config,
		// Name cannot be empty, and the length must be between 5 and 20.
		validation.Field(&config.StartDate, validation.Date("2006-01-02")),
		validation.Field(&config.EndDate, validation.Date("2006-01-02")),
		validation.Field(&config.Mode, validation.In("history", "regular", "regular_early")),
		validation.Field(&config.Source, validation.In("hits", "visits")),
		validation.Field(&config.Clickhouse),
	)
}

// parser.add_argument('-mode', help = 'Mode (one of [history, reqular, regular_early])')
// parser.add_argument('-source', help = 'Source (hits or visits)')
