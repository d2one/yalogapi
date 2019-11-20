package yalogapi

import (
	"fmt"
)

type YaLogApi struct {
	config *Config
}

type UserRequest struct {
	Token     string
	CounterID string
	StartDate string
	EndDate   string
	Source    string
	Fields    []string
}

func NewUserRequest(config *Config) *UserRequest {
	var fields []string
	switch config.Source {
	case "hits":
		fields = config.Logsapi.HitsField
		break
	case "visits":
		fields = config.Logsapi.VisitsField
		break
	}
	return &UserRequest{
		Token:     config.Logsapi.Token,
		CounterID: config.Logsapi.CounterID,
		StartDate: config.StartDate,
		EndDate:   config.EndDate,
		Source:    config.Source,
		Fields:    fields,
	}
}

func NewYaLogApi(config *Config) *YaLogApi {
	config.Init()
	return &YaLogApi{config: config}
}

func (yalogapi *YaLogApi) Run() {
	fmt.Println(yalogapi.config)
	fmt.Println(yalogapi.config.Clickhouse)
	fmt.Println(yalogapi.config.Logsapi)
	fmt.Println("YaLogApi")
	// userRequest := NewUserRequest(yalogapi.config)

}
