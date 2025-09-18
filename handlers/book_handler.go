package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/jerome-wilson/GO-REST-API/models"
	"github.com/jerome-wilson/GO-REST-API/storage"
)


func HandleBooks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		path := r.URL.Path
		parts := strings.Split(path, "/")

		if len(parts) == 2 || parts[2] == "" {
			respondWithJSON(w, storage.Books)
			return
		}

		idStr := parts[2]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "error - ID not found", http.StatusNotFound)
			return
		}

		for _, b := range storage.Books {
			if b.ID == id {
				respondWithJSON(w, b)
				return
			}
		}

		http.Error(w, "Book not found", http.StatusNotFound)

	case http.MethodPost:
		var newBook models.Book

		err := json.NewDecoder(r.Body).Decode(&newBook)
		if err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		newBook.ID = len(storage.Books) + 1
		storage.Books = append(storage.Books, newBook)

		w.WriteHeader(http.StatusCreated)
		respondWithJSON(w, newBook)

	case http.MethodPut:
		path := r.URL.Path
		parts := strings.Split(path, "/")

		if len(parts) < 3 || parts[2] == "" {
			http.Error(w, "ID must be provided", http.StatusBadRequest)
			return
		}

		idStr := parts[2]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid book ID", http.StatusBadRequest)
			return
		}

		var updatedBook models.Book

		err = json.NewDecoder(r.Body).Decode(&updatedBook)
		if err != nil {
			http.Error(w, "Invalid Request Body", http.StatusBadRequest)
			return
		}

		for i, b := range storage.Books {
			if b.ID == id {
				updatedBook.ID = id
				storage.Books[i] = updatedBook
				respondWithJSON(w, updatedBook)
				return
			}
		}

		http.Error(w, "Book not found", http.StatusNotFound)

	case http.MethodDelete:
		path := r.URL.Path
		parts := strings.Split(path, "/")

		if len(parts) < 3 || parts[2] == "" {
			http.Error(w, "ID must be provided", http.StatusBadRequest)
			return
		}

		idStr := parts[2]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid Book ID", http.StatusBadRequest)
			return
		}

		for i, b := range storage.Books {
			if b.ID == id {
				storage.Books = append(storage.Books[:i], storage.Books[i+1:]...)
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}
		http.Error(w, "Book Not Found", http.StatusNotFound)
	}
}

func respondWithJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
