package yalogapi

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
)

type ApiRequest struct {
	Requests []UserRequest
}

type Requests []UserRequest

const requestNewStatus = "new"

const requestCreatedStatus = "created"

const requestProcessedStatus = "processed"

const requestFailedStatus = "processing_failed"

const requestExpiredStatus = "cleaned_automatically_as_too_old"

const requestCleanedStatus = "cleaned_by_user"

const dateLayout = "2006-01-02"

const delayBetweenRequests = 20

func (yalogapi *YaLogApi) Run() {
	userRequest := NewUserRequest(yalogapi.config)
	requests, err := userRequest.getAPIRequests()

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, request := range requests {
		request.logCreate()

		err := request.create()
		if err != nil {
			request.logRequestError(err, "Create task")
			return
		}

		// @TODO use channel instead
		for request.TaskStatus.Status != requestProcessedStatus {
			// sleep
			// get status
			// exit if error
		}

		// save data
		// return for clickhous
	}
}

func (userRequest UserRequest) logCreate() {
	fmt.Println(
		fmt.Sprintf(
			"Create task. CounterID: %s, date preod: %s = %s",
			userRequest.CounterID,
			userRequest.StartDate,
			userRequest.EndDate,
		),
	)
}

func (userRequest UserRequest) logRequestError(err error, operation string) {
	fmt.Println(
		errors.Wrapf(
			err,
			"Operation %s. CounterID: %s, date preod: %s = %s",
			operation,
			userRequest.CounterID,
			userRequest.StartDate,
			userRequest.EndDate,
		),
	)
}

func (userRequest *UserRequest) create() error {
	taskLog, err := userRequest.createTask()

	if err != nil {
		return errors.Wrapf(
			err,
			"CounterID: %s, Date period: %s - %s",
			userRequest.CounterID,
			userRequest.StartDate,
			userRequest.EndDate,
		)
	}

	userRequest.RequestID = taskLog.RequestID
	userRequest.updateStatus(taskLog.Status)
	return nil
}

func (userRequest *UserRequest) check() error {
	taskStatus, err := userRequest.getStatus()

	if err != nil {
		return errors.Wrapf(
			err,
			"CounterID: %s, Date period: %s - %s",
			userRequest.CounterID,
			userRequest.StartDate,
			userRequest.EndDate,
		)
	}

	userRequest.TaskStatus = taskStatus
	return nil
}

func (userRequest UserRequest) getAPIRequests() (Requests, error) {
	evaluate, err := UserRequest.getEvaluation(userRequest)
	if err != nil {
		return nil, err
	}

	if !evaluate.Possible && evaluate.MaxPossibleDayQuantity == 0 {
		return nil, errors.New("Can not get data. max_possible_day_quantity = 0")
	}

	var requests Requests
	if evaluate.Possible {
		request := userRequest
		request.TaskStatus.Status = requestNewStatus
		requests = append(requests, request)
		return requests, nil
	}

	days, err := UserRequest.getDays(userRequest)

	if err != nil {
		return nil, err
	}

	requestsNumber := int(days/evaluate.MaxPossibleDayQuantity) + 1
	daysInPeriod := int(days/requestsNumber) + 1

	startDate := userRequest.StartDate
	endDate := userRequest.EndDate
	for i := 0; i < requestsNumber; i++ {
		startDate, endDate := getNewDates(startDate, endDate, daysInPeriod, i)
		userRequest.StartDate = startDate
		userRequest.EndDate = endDate

		request := userRequest
		request.TaskStatus.Status = requestNewStatus
		requests = append(requests, request)
	}

	return requests, nil
}

func (userRequest *UserRequest) updateStatus(status string) {
	if status == requestFailedStatus || status == requestExpiredStatus {
		userRequest.TaskStatus.Status = requestFailedStatus
	} else {
		userRequest.TaskStatus.Status = status
	}
}

func (userRequest UserRequest) getDays() (int, error) {
	from := userRequest.StartDate
	dateFrom, err := time.Parse(dateLayout, from)
	if err != nil {
		return 0, errors.New("Can not parse start date")
	}

	to := userRequest.EndDate
	dateTo, err := time.Parse(dateLayout, to)
	if err != nil {
		return 0, errors.New("Can not parse end date")
	}

	days := int(dateTo.Sub(dateFrom).Hours() / 24)

	return days, nil
}
