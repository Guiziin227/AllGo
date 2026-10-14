package task

import (
	"fmt"
	"strings"
	"time"
)

func (t *Task) String() string {
	status := " "
	if t.IsCompleted {
		status = "X"
	}
	return fmt.Sprintf("[%s] %d: %s", status, t.ID, t.Name)
}

type Task struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	IsCompleted bool      `json:"is_completed"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewTask(id int, name string) (*Task, error) {
	if name == "" || strings.TrimSpace(name) == "" {
		return nil, ErrTaskNameNotEmpty
	}

	return &Task{
		ID:        id,
		Name:      name,
		CreatedAt: time.Now(),
	}, nil
}

func (t *Task) Complete() error {
	if t.IsCompleted {
		return ErrTaskAlreadyClosed
	}
	t.IsCompleted = true
	return nil
}
