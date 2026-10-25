package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/guiziin227/allgo/biblioteca-api/library"
)

func main() {
	repo := library.NewMemoryRepository()
	service := library.NewService(repo)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /books", func(w http.ResponseWriter, r *http.Request) {

		// O motivo de usar withtimeout é para evitar vazamento de memória e de timers.
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()

		var input struct {
			ID         string `json:"id"`
			Title      string `json:"title"`
			Author     string `json:"author"`
			TotalStock int    `json:"total_stock"`
		}

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		book, err := service.CreateBook(ctx, input.ID, input.Title, input.Author, input.TotalStock)
		if errors.Is(err, library.ErrBookAlreadyExists) {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		if errors.Is(err, library.ErrInvalidBook) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(book)
		if errors.Is(err, context.DeadlineExceeded) {
			http.Error(w, "a operação demorou demais (timeout)", http.StatusGatewayTimeout)
			return
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	mux.HandleFunc("GET /books", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()

		books, err := service.ListBooks(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(books)
		if errors.Is(err, context.DeadlineExceeded) {
			http.Error(w, "a operação demorou demais (timeout)", http.StatusGatewayTimeout)
			return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	//mux.HandleFunc("GET /books", func(w http.ResponseWriter, r *http.Request) {
	//	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	//	defer cancel()
	//
	//	books, err := service.ListBooks(ctx)
	//	if err != nil {
	//		if errors.Is(err, context.DeadlineExceeded) {
	//			http.Error(w, "a operacao demorou demais (timeout)", http.StatusGatewayTimeout)
	//			return
	//		}
	//		http.Error(w, err.Error(), http.StatusInternalServerError)
	//		return
	//	}
	//
	//	w.Header().Set("Content-Type", "application/json")
	//	w.WriteHeader(http.StatusOK)
	//	json.NewEncoder(w).Encode(books)
	//})

	mux.HandleFunc("POST /books/{id}/borrow", func(w http.ResponseWriter, r *http.Request) {

		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()

		id := r.PathValue("id")
		book, err := service.BorrowBook(ctx, id)

		if errors.Is(err, library.ErrBookNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		if errors.Is(err, library.ErrOutOfStock) {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	mux.HandleFunc("POST /books/{id}/return", func(w http.ResponseWriter, r *http.Request) {

		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()

		id := r.PathValue("id")

		book, err := service.ReturnBook(ctx, id)

		if errors.Is(err, library.ErrBookNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		if errors.Is(err, library.ErrStockFull) {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(book)
		if errors.Is(err, context.DeadlineExceeded) {
			http.Error(w, "a operação demorou demais (timeout)", http.StatusGatewayTimeout)
			return
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	fmt.Println("Listening on :8000")
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal(err)
		return
	}

}
