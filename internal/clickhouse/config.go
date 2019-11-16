package clickhouse

import (
	"fmt"
)

type ClickhouseConfig struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	VisitsTable string `mapstructure:"visitsTable"`
	HitsTable   string `mapstructure:"hitsTable"`
	Database    string `mapstructure:"database"`
}

func NewClickhouseConfig() (*ClickhouseConfig, error) {
	return &ClickhouseConfig{}, nil
}

func (config *ClickhouseConfig) getDSN() string {
	return fmt.Sprintf("tcp://%s:%d?username=%s&password=%s&database=%s&read_timeout=10&write_timeout=20",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.Database,
	)
}
