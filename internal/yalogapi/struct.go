package yalogapi

import "github.com/d2one/yalogapi/internal/clickhouse"

type YaLogApi struct {
	config     *Config
	clickhouse *clickhouse.Clickhouse
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

type TaskLog struct {
	RequestID int    `json:"request_id"`
	CounterID int    `json:"counter_id"`
	Status    string `json:"status"`
}

type CreateTaskResponse struct {
	TaskLog TaskLog `json:"log_request"`
}

type LogPart struct {
	PartNumber int `json:"part_number"`
	Size       int `json:"size"`
}

type TaskStatus struct {
	Status string `json:"status"`
	Parts  []LogPart
}

type GetStatusResponse struct {
	TaskStatus TaskStatus `json:"log_request"`
}
