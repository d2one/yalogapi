package yalogapi

import (
	"errors"
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

type LogRequestEvaluation struct {
	Possible               bool `json:"possible"`
	MaxPossibleDayQuantity int  `json:"max_possible_day_quantity"`
}

type EvaluateResponse struct {
	LogRequestEvaluation LogRequestEvaluation `json:"log_request_evaluation"`
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

func getEvaluation(request UserRequest) (LogRequestEvaluation, error) {
	return LogRequestEvaluation{true, 20}, nil
}

func GetApiRequests(request UserRequest) ([]UserRequest, error) {
	evaluate, err := getEvaluation(request)

	if err != nil {
		return nil, err
	}

	if !evaluate.Possible && evaluate.MaxPossibleDayQuantity == 0 {
		return nil, errors.New("Can not get data. max_possible_day_quantity = 0")
	}

	requests := []UserRequest{}
	if evaluate.Possible {
		requests = append(requests, request)
		return requests, nil
	}

	// build new request here

	return requests, nil
}
