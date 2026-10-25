package library

import (
	"context"
	"errors"
	"sync"
)

type Service struct {
	mu   sync.Mutex
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateBook(ctx context.Context, id, title, author string, totalStock int) (*Book, error) {

	book, err := NewBook(id, title, author, totalStock)
	if err != nil {
		return nil, err
	}

	_, err = s.repo.FindByID(ctx, book.ID)
	if err == nil {
		return nil, ErrBookAlreadyExists
	}
	if !errors.Is(err, ErrBookNotFound) {
		return nil, err
	}

	err = s.repo.Save(ctx, book)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (s *Service) ListBooks(ctx context.Context) ([]*Book, error) {
	return s.repo.FindAll(ctx)
}

func (s *Service) BorrowBook(ctx context.Context, id string) (*Book, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	book, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	err = book.Borrow()
	if err != nil {
		return nil, err
	}

	err = s.repo.Save(ctx, book)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (s *Service) ReturnBook(ctx context.Context, id string) (*Book, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	book, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	err = book.Return()
	if err != nil {
		return nil, err
	}

	err = s.repo.Save(ctx, book)
	if err != nil {
		return nil, err
	}

	return book, nil
}
