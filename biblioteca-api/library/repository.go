package library

import "sync"

type Repository interface {
	Save(book *Book) error
	FindByID(id string) (*Book, error)
	FindAll() ([]*Book, error)
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

func (r *MemoryRepository) Save(book *Book) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.books[book.ID] = book
	return nil
}

func (r *MemoryRepository) FindByID(id string) (*Book, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	book, exists := r.books[id]
	if !exists {
		return nil, ErrBookNotFound
	}
	return book, nil
}

func (r *MemoryRepository) FindAll() ([]*Book, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	books := make([]*Book, 0, len(r.books))
	for _, book := range r.books {
		books = append(books, book)
	}
	return books, nil
}
