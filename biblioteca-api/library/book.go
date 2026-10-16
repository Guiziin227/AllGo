package library

import (
	"errors"
	"strings"
)

var (
	ErrInvalidBook       = errors.New("invalid book")
	ErrOutOfStock        = errors.New("book is out of stock")
	ErrStockFull         = errors.New("book stock is full")
	ErrBookNotFound      = errors.New("book not found")
	ErrBookAlreadyExists = errors.New("book already exists")
)

type Book struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	TotalStock int    `json:"total_stock"`
	Available  int    `json:"available"`
}

func NewBook(id, title, author string, totalStock int) (*Book, error) {

	id = strings.TrimSpace(id)
	title = strings.TrimSpace(title)
	author = strings.TrimSpace(author)

	if id == "" || title == "" || author == "" || totalStock <= 0 {
		return nil, ErrInvalidBook
	}
	return &Book{
		ID:         id,
		Title:      title,
		Author:     author,
		TotalStock: totalStock,
		Available:  totalStock,
	}, nil
}

func (b *Book) Borrow() error {

	if b.Available <= 0 {
		return ErrOutOfStock
	}

	b.Available--

	return nil
}

func (b *Book) Return() error {
	if b.Available >= b.TotalStock {
		return ErrStockFull
	}
	b.Available++
	return nil
}
