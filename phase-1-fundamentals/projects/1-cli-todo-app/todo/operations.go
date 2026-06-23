package todo

import (
	"fmt"
	"time"
)

func CreateTodo(id int, name string, task string, priority TodoPriorityType) (*TodoItem, error) {
	if name == "" {
		return nil, ErrEmptyName
	}
	if task == "" {
		return nil, ErrEmptyTask
	}
	_, valid := PriorityLabel[priority]
	if !valid {
		return nil, ErrInvalidPriority
	}

	return &TodoItem{
		Id:        id,
		Name:      name,
		Task:      task,
		Priority:  priority,
		Status:    PENDING,
		CreatedAt: time.Now(),
	}, nil
}

func (t *TodoItem) ChangePriority(newPriority TodoPriorityType) error {
	_, valid := PriorityLabel[newPriority]
	if !valid {
		return fmt.Errorf("ChangePriority: %w", ErrInvalidPriority)
	}
	t.Priority = newPriority
	return nil
}

func (t *TodoItem) MarkCompleted() error {
	if t.Status == COMPLETED {
		return fmt.Errorf("MarkCompleted (id=%d): %w", t.Id, ErrAlreadyCompleted)
	}
	if t.Status == EXPIRED {
		return fmt.Errorf("MarkCompleted (id=%d): %w", t.Id, ErrAlreadyExpired)
	}
	t.Status = COMPLETED
	t.CompletedAt = time.Now()
	return nil
}

func (t *TodoItem) MarkExpired() error {
	if t.Status == COMPLETED {
		return fmt.Errorf("MarkExpired (id=%d): %w", t.Id, ErrAlreadyCompleted)
	}
	if t.Status == EXPIRED {
		return fmt.Errorf("MarkExpired (id=%d): %w", t.Id, ErrAlreadyExpired)
	}
	t.Status = EXPIRED
	return nil
}

func (t *TodoItem) AddExpiry(expiryDate time.Time) error {
	if expiryDate.Before(time.Now()) {
		return fmt.Errorf("AddExpiry (id=%d): %w", t.Id, ErrExpiryInPast)
	}
	t.ExpiryAt = expiryDate.Local()
	return nil
}