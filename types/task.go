package types

import "time"

type Status string

const (
	WAITING Status = "Waiting"
	WORKING Status = "Working"
	WRITING Status = "Waiting"
)

type Task struct {
	ID         int       `json:"id"`
	Payload    Payload   `json:"payload"`
	Result     Result    `json:"result"`
	Status     Status    `json:"status"`
	PlacedAt   time.Time `json:"placedAt"`
	StartedAt  time.Time `json:"startedAt"`
	FinishedAt time.Time `json:"finishedAt"`
}

func NewTask(payload Payload, status Status, placedAt time.Time) *Task {
	return &Task{
		Payload:  payload,
		Status:   status,
		PlacedAt: placedAt,
	}
}
