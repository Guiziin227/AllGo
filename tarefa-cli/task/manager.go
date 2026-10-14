package task

import (
	"encoding/json"
	"errors"
	"os"
)

var (
	ErrTaskNotFound      = errors.New("task not found")
	ErrTaskAlreadyClosed = errors.New("task is already completed")
	ErrTaskNameNotEmpty  = errors.New("task name cannot be empty")
)

type Manager struct {
	nextID int
	tasks  []*Task
}

func NewManager() *Manager {
	return &Manager{
		nextID: 1,
		tasks:  []*Task{},
	}
}

func (m *Manager) Add(name string) (*Task, error) {

	t, err := NewTask(m.nextID, name)
	if err != nil {
		return nil, err
	}
	m.tasks = append(m.tasks, t)
	m.nextID++
	return t, nil
}

func (m *Manager) List() []*Task {
	return m.tasks
}

func (m *Manager) Complete(id int) error {
	for _, t := range m.tasks {
		if id == t.ID {
			return t.Complete()
		}
	}
	return ErrTaskNotFound
}

func (m *Manager) Save(filename string) error {
	data, err := json.MarshalIndent(m.tasks, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (m *Manager) Load(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	var t []*Task
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	m.tasks = t
	m.nextID = 1
	for _, task := range t {
		if task.ID >= m.nextID {
			m.nextID = task.ID + 1
		}
	}

	return nil
}
