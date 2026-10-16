package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/guiziin227/allgo/biblioteca-api/library"
)

func main() {

	repo := library.NewMemoryRepository()
	service := library.NewService(repo)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /books", func(w http.ResponseWriter, r *http.Request) {
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

		book, err := service.CreateBook(input.ID, input.Title, input.Author, input.TotalStock)
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
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	mux.HandleFunc("GET /books", func(w http.ResponseWriter, r *http.Request) {
		books, err := service.ListBooks()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(books)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	mux.HandleFunc("POST /books/{id}/borrow", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		book, err := service.BorrowBook(id)

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
		id := r.PathValue("id")

		book, err := service.ReturnBook(id)

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
