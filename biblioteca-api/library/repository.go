package library

import (
	"context"
	"sync"
)

type Repository interface {
	Save(ctx context.Context, book *Book) error
	FindByID(ctx context.Context, id string) (*Book, error)
	FindAll(ctx context.Context) ([]*Book, error)
}

type MemoryRepository struct {
	mu    sync.RWMutex
	books map[string]*Book
}

var _ Repository = (*MemoryRepository)(nil)

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		books: make(map[string]*Book),
	}
}

func (r *MemoryRepository) Save(ctx context.Context, book *Book) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	r.books[book.ID] = book
	return nil
}

func (r *MemoryRepository) FindByID(ctx context.Context, id string) (*Book, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	r.mu.RLock()
	defer r.mu.RUnlock()
	book, exists := r.books[id]
	if !exists {
		return nil, ErrBookNotFound
	}
	return book, nil
}

func (r *MemoryRepository) FindAll(ctx context.Context) ([]*Book, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	r.mu.RLock()
	defer r.mu.RUnlock()
	books := make([]*Book, 0, len(r.books))
	for _, book := range r.books {
		books = append(books, book)
	}
	return books, nil
}
