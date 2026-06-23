package todo

import (
	"errors"
	"time"
)

var ErrTodoNotFound = errors.New("todo not found")
var ErrAlreadyCompleted = errors.New("todo is already completed")
var ErrAlreadyExpired = errors.New("todo is already expired")
var ErrInvalidPriority = errors.New("invalid priority — use: low, medium, high")
var ErrEmptyName = errors.New("todo name cannot be empty")
var ErrEmptyTask = errors.New("todo task cannot be empty")
var ErrExpiryInPast = errors.New("expiry date cannot be in the past")

type TodoStatusType string
type TodoPriorityType int8

const (
	LOW_PRIORITY    TodoPriorityType = iota + 1 // 1
	MEDIUM_PRIORITY                             // 2
	HIGH_PRIORITY                               // 3
)

var PriorityLabel = map[TodoPriorityType]string{
	LOW_PRIORITY:    "low",
	MEDIUM_PRIORITY: "medium",
	HIGH_PRIORITY:   "high",
}

var PriorityFromString = map[string]TodoPriorityType{
	"low":    LOW_PRIORITY,
	"medium": MEDIUM_PRIORITY,
	"high":   HIGH_PRIORITY,
}

const (
	PENDING   TodoStatusType = "pending"
	COMPLETED TodoStatusType = "completed"
	EXPIRED   TodoStatusType = "expired"
)

type TodoItem struct {
	Id          int              `json:"id"`
	Name        string           `json:"name"`
	Task        string           `json:"task"`
	Status      TodoStatusType   `json:"status"`
	Priority    TodoPriorityType `json:"priority"`
	ExpiryAt    time.Time        `json:"expiry_at,omitempty"`
	CreatedAt   time.Time        `json:"created_at"`
	CompletedAt time.Time        `json:"completed_at,omitempty"`
}
