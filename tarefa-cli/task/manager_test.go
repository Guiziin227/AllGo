package task

import (
	"errors"
	"testing"
)

func TestManagerAddAndComplete(t *testing.T) {
	m := NewManager()

	// 1. Testa adicionar tarefa válida
	created, err := m.Add("Estudar Go")
	if err != nil {
		t.Fatalf("não esperava erro ao criar tarefa, mas recebeu: %v", err)
	}
	if created.ID != 1 {
		t.Errorf("esperava ID 1, mas recebeu %d", created.ID)
	}

	// 2. Testa concluir a tarefa
	if err := m.Complete(1); err != nil {
		t.Errorf("erro inesperado ao completar tarefa: %v", err)
	}
	if !created.IsCompleted {
		t.Error("a tarefa deveria estar marcada como concluída")
	}
}

func TestManagerAddEmptyName(t *testing.T) {
	m := NewManager()

	_, err := m.Add(" ")
	if !errors.Is(err, ErrTaskNameNotEmpty) {
		t.Errorf("esperava erro %v, mas recebeu: %v", ErrTaskNameNotEmpty, err)
	}
}

func TestManagerCompleteError(t *testing.T) {
	m := NewManager()

	if err := m.Complete(1); !errors.Is(err, ErrTaskNotFound) {
		t.Errorf("esperava erro %v, mas recebeu: %v", ErrTaskNotFound, err)
	}
}
