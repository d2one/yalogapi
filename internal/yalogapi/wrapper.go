package yalogapi

import (
	"time"

	"github.com/pkg/errors"
)

type ApiRequest struct {
	Requests []UserRequest
}

type RequestWrapper struct {
	Status  string
	Request UserRequest
}

type Requests []RequestWrapper

const requestNewStatus = "new"

const requestFinishedStatus = "succeed"

const requestFailedStatus = "failed"

const dateLayout = "2006-01-02"

func (userRequest UserRequest) GetAPIRequests() (Requests, error) {
	evaluate, err := UserRequest.getEvaluation(userRequest)
	if err != nil {
		return nil, err
	}

	if !evaluate.Possible && evaluate.MaxPossibleDayQuantity == 0 {
		return nil, errors.New("Can not get data. max_possible_day_quantity = 0")
	}

	var requests Requests;
	if evaluate.Possible {
		requestWrapper := RequestWrapper{}
		requestWrapper.Request = userRequest
		requestWrapper.Status = requestNewStatus
		requests = append(requests, requestWrapper)
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

		requestWrapper := RequestWrapper{}
		requestWrapper.Request = userRequest
		requestWrapper.Status = requestNewStatus
		requests = append(requests, requestWrapper)
	}

	return requests, nil
}

func (requestWrapper *RequestWrapper) updateStatus(status TaskStatus) {
	if status.Status == requestFinishedStatus {
		requestWrapper.Status = requestFinishedStatus
	}
	if status.Status == requestFailedStatus {
		requestWrapper.Status = requestFailedStatus
	}
}

func (request UserRequest) getDays() (int, error) {
	from := request.StartDate
	dateFrom, err := time.Parse(dateLayout, from)
	if err != nil {
		return 0, errors.New("Can not parse start date")
	}

	to := request.EndDate
	dateTo, err := time.Parse(dateLayout, to)
	if err != nil {
		return 0, errors.New("Can not parse end date")
	}

	days := int(dateTo.Sub(dateFrom).Hours() / 24)

	return days, nil
}
