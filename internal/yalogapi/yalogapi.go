package yalogapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/d2one/yalogapi/internal/clickhouse"
	"github.com/pkg/errors"
)

// NewYaLogApi create YaLogApi
func NewYaLogApi(config *Config) *YaLogApi {
	config.Init()
	return &YaLogApi{config: config, clickhouse: clickhouse.NewClickhouse(config.Clickhouse)}
}

func (yalogapi *YaLogApi) Run() {
	fmt.Println(yalogapi.config.Types)
	// userRequest := NewUserRequest(yalogapi.config)

	// result, err := createTask(userRequest)
	// if err != nil {
	// 	fmt.Println("error when do request")
	// }

	// status, partCount, err := getStatus(result, userRequest.Token)

	// fmt.Println(status)
	// fmt.Println(partCount)
}

const host string = "https://api-metrika.yandex.ru"

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

// getEvaluation returns estimation of Logs API
func getEvaluation(userRequest UserRequest) (LogRequestEvaluation, error) {
	uri := fmt.Sprintf(
		"%s/management/v1/counter/%s/logrequests/evaluate?",
		host,
		userRequest.CounterID,
	)

	query := createUserQuery(userRequest)
	response, err := doRequest("GET", uri, userRequest.Token, query)
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

func createUserQuery(userRequest UserRequest) url.Values {
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

func doRequest(method string, uri string, token string, query url.Values) ([]byte, error) {
	request, err := http.NewRequest(method, uri, nil)

	errorMessage := fmt.Sprintf("Yandex api request failed. Uri %s", uri)

	if err != nil {
		return nil, errors.Wrapf(err, errorMessage)
	}

	request.Header.Add(
		"Authorization",
		fmt.Sprintf("OAuth %s", token),
	)

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

func getAPIRequests(request UserRequest) ([]UserRequest, error) {
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

	days, err := getDays(request)
	fmt.Println(days)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	requestsNumber := int(days/evaluate.MaxPossibleDayQuantity) + 1
	daysInPeriod := int(days/requestsNumber) + 1

	startDate := request.StartDate
	endDate := request.EndDate
	for i := 0; i < requestsNumber; i++ {
		startDate, endDate := getNewDates(startDate, endDate, daysInPeriod, i)
		request.StartDate = startDate
		request.EndDate = endDate

		// @TODO mark every request (append a new struct with request body and request status)
		requests = append(requests, request)
	}

	return requests, nil
}

// @TODO don`t validate date here
func getDays(request UserRequest) (int, error) {
	from := request.StartDate
	dateFrom, err := time.Parse("2006-01-02", from)
	if err != nil {
		return 0, errors.New("Can not parse start date")
	}

	to := request.EndDate
	dateTo, err := time.Parse("2006-01-02", to)
	if err != nil {
		return 0, errors.New("Can not parse end date")
	}

	days := int(dateTo.Sub(dateFrom).Hours() / 24)

	return days, nil
}

// @TODO don`t validate date here
func getNewDates(from string, to string, daysInPeriod int, partNumber int) (string, string) {
	dateFrom, _ := time.Parse("2006-01-02", from)
	newDateFrom := dateFrom.AddDate(0, 0, partNumber*daysInPeriod)

	dateTo, _ := time.Parse("2006-01-02", to)
	newDateTo := dateFrom.AddDate(0, 0, (partNumber+1)*daysInPeriod-1)

	if newDateTo.After(dateTo) {
		newDateTo = dateTo
	}

	return newDateFrom.Format("2006-01-02"), newDateTo.Format("2006-01-02")
}

// createTask method creates a Logs API task to generate data
func createTask(userRequest UserRequest) (TaskLog, error) {
	uri := fmt.Sprintf(
		"%s/management/v1/counter/%s/logrequests/?",
		host,
		userRequest.CounterID,
	)

	query := createUserQuery(userRequest)
	response, err := doRequest("POST", uri, userRequest.Token, query)

	сreateTaskResponse := CreateTaskResponse{}
	err = json.Unmarshal(response, &сreateTaskResponse)
	if err != nil {
		return TaskLog{}, err
	}

	return сreateTaskResponse.TaskLog, nil
}

// getStatus returns current tasks status and part counts
func getStatus(taskLog TaskLog, token string) (string, int, error) {
	uri := fmt.Sprintf(
		"%s/management/v1/counter/%d/logrequest/%d",
		host,
		taskLog.CounterID,
		taskLog.RequestID,
	)

	query := url.Values{}
	response, err := doRequest("GET", uri, token, query)

	getStatusResponse := GetStatusResponse{}
	err = json.Unmarshal(response, &getStatusResponse)
	if err != nil {
		return "", 0, err
	}

	return getStatusResponse.TaskStatus.Status, len(getStatusResponse.TaskStatus.Parts), nil
}
