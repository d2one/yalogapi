package yalogapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/d2one/yalogapi/internal/clickhouse"
	"github.com/pkg/errors"
)

// NewYaLogApi create YaLogApi
func NewYaLogApi(config *Config) *YaLogApi {
	config.Init()
	clickhouse, err := clickhouse.NewClickhouse(config.Clickhouse)
	if err != nil {
		fmt.Println(err)
	}
	return &YaLogApi{config: config, clickhouse: clickhouse}
}

// clickhouse.save_data(api_request.user_request.source,
// 	api_request.user_request.fields,
// 	output_data)

// NewUserRequest create UserRequest
func NewUserRequest(config *Config) UserRequest {
	var fields []string
	switch config.Source {
	case "hits":
		fields = config.Logsapi.HitsField
		break
	case "visits":
		fields = config.Logsapi.VisitsField
		break
	}

	return UserRequest{
		Token:     config.Logsapi.Token,
		CounterID: config.Logsapi.CounterID,
		StartDate: config.StartDate,
		EndDate:   config.EndDate,
		Source:    config.Source,
		Fields:    fields,
	}
}

func (yalogapi *YaLogApi) Run() {
	// fmt.Println(yalogapi.config.Types)
	// yalogapi.clickhouse.Check(yalogapi.config.Source, yalogapi.config.getMappedFilds())

	//userRequest := NewUserRequest(yalogapi.config)

	// fmt.Println(requests)

	// taskLog, err := UserRequest.createTask(userRequest)
	// if err != nil {
	// 	fmt.Println("error when do request")
	// }

	// taskStatus, err := UserRequest.getStatus(userRequest, taskLog)

	// fmt.Println(taskStatus)
}

const host string = "https://api-metrika.yandex.ru"

func (userRequest UserRequest) getEvaluation() (LogRequestEvaluation, error) {
	uri := fmt.Sprintf(
		"%s/management/v1/counter/%s/logrequests/evaluate?",
		host,
		userRequest.CounterID,
	)

	query := UserRequest.buildQuery(userRequest)

	httpRequest := httpRequest{
		"GET",
		uri,
		userRequest.Token,
		query,
	}

	response, err := doRequest(httpRequest)
	if err != nil {
		return LogRequestEvaluation{}, err
	}

	evaluateResponse := EvaluateResponse{}
	err = json.Unmarshal(response, &evaluateResponse)
	if err != nil {
		return LogRequestEvaluation{}, errors.Wrap(err, "error when Unmarshal")
	}

	return evaluateResponse.LogRequestEvaluation, nil
}

// createTask method creates a Logs API task to generate data
func (userRequest UserRequest) createTask() (TaskLog, error) {
	uri := fmt.Sprintf(
		"%s/management/v1/counter/%s/logrequests/?",
		host,
		userRequest.CounterID,
	)

	query := UserRequest.buildQuery(userRequest)

	httpRequest := httpRequest{
		"POST",
		uri,
		userRequest.Token,
		query,
	}

	response, err := doRequest(httpRequest)

	сreateTaskResponse := CreateTaskResponse{}
	err = json.Unmarshal(response, &сreateTaskResponse)
	if err != nil {
		return TaskLog{}, err
	}

	return сreateTaskResponse.TaskLog, nil
}

// getStatus returns current tasks status and part counts
func (userRequest UserRequest) getStatus() (TaskStatus, error) {
	if userRequest.RequestID == 0 {
		err := errors.New("RequestID сan not be 0")
		return TaskStatus{}, err
	}

	uri := fmt.Sprintf(
		"%s/management/v1/counter/%s/logrequest/%d",
		host,
		userRequest.CounterID,
		userRequest.RequestID,
	)

	query := url.Values{}
	httpRequest := httpRequest{
		"GET",
		uri,
		userRequest.Token,
		query,
	}

	response, err := doRequest(httpRequest)

	getStatusResponse := GetStatusResponse{}
	err = json.Unmarshal(response, &getStatusResponse)
	if err != nil {
		return TaskStatus{}, err
	}

	return getStatusResponse.TaskStatus, nil
}

func (userRequest UserRequest) buildQuery() url.Values {
	query := url.Values{}

	query.Add("date1", userRequest.StartDate)
	query.Add("date2", userRequest.EndDate)

	if len(userRequest.Source) > 0 {
		query.Add("source", userRequest.Source)
	}

	if len(userRequest.Fields) > 0 {
		query.Add("fields", strings.Join(userRequest.Fields, ","))
	}

	return query
}

// @TODO check and return struct here
func (userRequest UserRequest) getReadyData(part int) ([]byte, error) {
	if userRequest.RequestID == 0 {
		err := errors.New("RequestID сan not be 0")
		return []byte(""), err
	}

	uri := fmt.Sprintf(
		"%s/management/v1/counter/%s/logrequest/%d/part/%d/download",
		host,
		userRequest.CounterID,
		userRequest.RequestID,
		part,
	)

	query := url.Values{}
	httpRequest := httpRequest{
		"GET",
		uri,
		userRequest.Token,
		query,
	}

	return doRequest(httpRequest)
}

// @TODO check and return struct here
func (userRequest UserRequest) clean() ([]byte, error) {
	uri := fmt.Sprintf(
		"%s/management/v1/counter/%s/logrequest/%d/clean",
		host,
		userRequest.CounterID,
		userRequest.RequestID,
	)

	query := url.Values{}
	httpRequest := httpRequest{
		"POST",
		uri,
		userRequest.Token,
		query,
	}

	return doRequest(httpRequest)
}
