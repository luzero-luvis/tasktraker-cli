package internal

import (
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CompletedAt time.Time `json:"CompletedAt"`
}
