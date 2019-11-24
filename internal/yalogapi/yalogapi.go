package yalogapi

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strings"
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

func NewYaLogApi(config *Config) *YaLogApi {
	config.Init()
	return &YaLogApi{config: config}
}

func (yalogapi *YaLogApi) Run() {
	userRequest := NewUserRequest(yalogapi.config)

	apiRequests, err := GetApiRequests(userRequest)
	if err != nil {
		fmt.Println("error when get evaluation from api")
	}

	fmt.Println(apiRequests)
}

// getEvaluation returns estimation of Logs API
func getEvaluation(userRequest *UserRequest) (LogRequestEvaluation, error) {
	url := fmt.Sprintf("%s/management/v1/counter/%s/logrequests/evaluate?", Host, userRequest.CounterID)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LogRequestEvaluation{}, errors.Wrap(err, "error when create new request")
	}

	request.Header.Add("Authorization", "OAuth "+userRequest.Token)

	query := request.URL.Query()
	query.Add("date1", userRequest.StartDate)
	query.Add("date2", userRequest.EndDate)
	query.Add("source", userRequest.Source)
	query.Add("fields", strings.Join(userRequest.Fields, ","))

	request.URL.RawQuery = query.Encode()

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return LogRequestEvaluation{}, errors.Wrap(err, "error when do request")
	}

	result, error := ioutil.ReadAll(response.Body)
	if error != nil {
		return LogRequestEvaluation{}, errors.Wrap(err, "error when read response")
	}

	logRequest := EvaluateResponse{}
	err = json.Unmarshal(result, &logRequest)
	if err != nil {
		return LogRequestEvaluation{}, errors.Wrap(err, "error when Unmarshal")
	}

	return logRequest.LogRequestEvaluation, nil
}

func GetApiRequests(request *UserRequest) ([]UserRequest, error) {
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
