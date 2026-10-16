package encurtador

import (
	"errors"
	"math/rand"
	"strings"
	"sync"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var (
	ErrEmptyURL     = errors.New("url is empty")
	ErrLinkNotFound = errors.New("url not found")
)

type Link struct {
	Code        string    `json:"code"`
	OriginalURL string    `json:"original_url"`
	Clicks      int       `json:"clicks"`
	CreatedAt   time.Time `json:"created_at"`
}

type Shortener struct {
	mu    sync.RWMutex
	links map[string]*Link
}

func NewShortener() *Shortener {
	return &Shortener{
		links: make(map[string]*Link),
	}
}

func (s *Shortener) Shorten(originalURL string) (*Link, error) {

	url := strings.TrimSpace(originalURL)
	if url == "" {
		return nil, ErrEmptyURL
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	var code string
	for {
		code = generateCode(6)
		if _, exists := s.links[code]; !exists {
			break
		}
	}

	link := &Link{
		Code:        code,
		OriginalURL: url,
		CreatedAt:   time.Now(),
	}
	s.links[code] = link

	return link, nil
}

func (s *Shortener) Resolve(code string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	link, ok := s.links[code]
	if !ok {
		return "", ErrLinkNotFound
	}
	link.Clicks++
	return link.OriginalURL, nil
}

func (s *Shortener) Stats(code string) (Link, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	link, ok := s.links[code]
	if !ok {
		return Link{}, ErrLinkNotFound
	}
	return *link, nil
}

func generateCode(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
