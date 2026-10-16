package main

import (
	"encoding/json"

	"fmt"
	"net/http"

	"github.com/guiziin227/allgo/encurtador/encurtador"
)

func main() {

	encurt := encurtador.NewShortener()
	mux := http.NewServeMux()

	mux.HandleFunc("POST /shorten", func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			URL string `json:"url"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		link, err := encurt.Shorten(input.URL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(link)
		if err != nil {
			return
		}
	})

	mux.HandleFunc("GET /{code}", func(w http.ResponseWriter, r *http.Request) {
		code := r.PathValue("code")

		link, err := encurt.Resolve(code)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Redirect(w, r, link, http.StatusTemporaryRedirect)
	})

	mux.HandleFunc("GET /stats/{code}", func(w http.ResponseWriter, r *http.Request) {
		code := r.PathValue("code")

		link, err := encurt.Stats(code)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(link)
		if err != nil {
			return
		}
	})

	fmt.Println("Servidor rodando em http://localhost:8080")
	if err := http.ListenAndServe(":8000", mux); err != nil {
		fmt.Println("Erro ao iniciar servidor:", err)
	}
}
