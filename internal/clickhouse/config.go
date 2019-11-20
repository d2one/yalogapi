package clickhouse

import (
	"fmt"

	"github.com/go-ozzo/ozzo-validation/is"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Config struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	VisitsTable string `mapstructure:"visitsTable"`
	HitsTable   string `mapstructure:"hitsTable"`
	Database    string `mapstructure:"database"`
}

func NewClickhouseConfig() (*Config, error) {
	return &Config{}, nil
}

func (config *Config) getDSN() string {
	return fmt.Sprintf("tcp://%s:%d?username=%s&password=%s&database=%s&read_timeout=10&write_timeout=20",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.Database,
	)
}

func (config Config) Validate() error {
	return validation.ValidateStruct(&config,
		// Name cannot be empty, and the length must be between 5 and 20.
		validation.Field(&config.Host, validation.Required, is.Host),
		validation.Field(&config.Port, validation.Required),
		validation.Field(&config.User, validation.Required),
		validation.Field(&config.Password, validation.Required),
		validation.Field(&config.Database, validation.Required),
	)
}
