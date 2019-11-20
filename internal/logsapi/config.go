package logsapi

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Config struct {
	Token       string   `mapstructure:"token"`
	CounterID   string   `mapstructure:"counter_id"`
	VisitsField []string `mapstructure:"visits_fields"`
	HitsField   []string `mapstructure:"hits_fields"`
}

func (config Config) Validate() error {
	return validation.ValidateStruct(&config,
		// Name cannot be empty, and the length must be between 5 and 20.
		validation.Field(&config.Token, validation.Required),
		validation.Field(&config.CounterID, validation.Required),
		validation.Field(&config.VisitsField, validation.Required),
		validation.Field(&config.HitsField, validation.Required),
	)
}
