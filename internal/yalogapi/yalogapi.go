package yalogapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/d2one/yalogapi/internal/clickhouse"
	"github.com/pkg/errors"
)

type YaLogApi struct {
	config     *Config
	clickhouse *clickhouse.Clickhouse
}

func NewYaLogApi(config *Config) *YaLogApi {
	config.Init()
	return &YaLogApi{config: config, clickhouse: clickhouse.NewClickhouse(config.Clickhouse)}
}

func (yalogapi *YaLogApi) Run() {
	fmt.Println(yalogapi.config.Types)
	// userRequest := NewUserRequest(yalogapi.config)

	// apiRequests, err := getAPIRequests(userRequest)
	// if err != nil {
	// 	fmt.Println("error when get evaluation from api")
	// }

	// fmt.Println(apiRequests)
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

const Host string = "https://api-metrika.yandex.ru"

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

// getEvaluation returns estimation of Logs API
func getEvaluation(userRequest *UserRequest) (LogRequestEvaluation, error) {
	uri := fmt.Sprintf(
		"%s/management/v1/counter/%s/logrequests/evaluate?",
		Host,
		userRequest.CounterID,
	)

	response, err := doGetRequest(uri, userRequest)
	if err != nil {
		return LogRequestEvaluation{}, err
	}

	logRequest := EvaluateResponse{}
	err = json.Unmarshal(response, &logRequest)
	if err != nil {
		return LogRequestEvaluation{}, errors.Wrap(err, "error when Unmarshal")
	}

	return logRequest.LogRequestEvaluation, nil
}

func doGetRequest(uri string, userRequest *UserRequest) ([]byte, error) {
	request, err := http.NewRequest("GET", uri, nil)

	errorMessage := fmt.Sprintf("Yandex api request failed. Uri %s", uri)

	if err != nil {
		return nil, errors.Wrapf(err, errorMessage)
	}

	request.Header.Add(
		"Authorization",
		fmt.Sprintf("OAuth %s", userRequest.Token),
	)

	query := request.URL.Query()

	query.Add("date1", userRequest.StartDate)
	query.Add("date2", userRequest.EndDate)

	if len(userRequest.Source) > 0 {
		query.Add("source", userRequest.Source)
	}

	if len(userRequest.Fields) > 0 {
		query.Add("fields", strings.Join(userRequest.Fields, ","))
	}
	request.URL.RawQuery = query.Encode()

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return nil, errors.Wrapf(err, errorMessage)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, errors.Wrapf(err, errorMessage, uri)
	}

	return body, nil
}

func getAPIRequests(request *UserRequest) ([]UserRequest, error) {
	evaluate, err := getEvaluation(request)
	if err != nil {
		return nil, err
	}

	if !evaluate.Possible && evaluate.MaxPossibleDayQuantity == 0 {
		return nil, errors.New("Can not get data. max_possible_day_quantity = 0")
	}

	requests := []UserRequest{}
	if evaluate.Possible {
		config := UserRequest{
			Token:     request.Token,
			CounterID: request.CounterID,
			StartDate: request.StartDate,
			EndDate:   request.EndDate,
			Source:    request.Source,
			Fields:    request.Fields,
		}
		requests = append(requests, config)
		return requests, nil
	}

	// build new request here

	return requests, nil
}
