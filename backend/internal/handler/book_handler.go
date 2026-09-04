package handler

import (
	"encoding/json"
	"io"
	"library-system/internal/model"
	"library-system/internal/service"
	"net/http"
	"strings"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var newBook model.Book

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&newBook)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
		return
	}

	err = decoder.Decode(&struct{}{})
	if err != io.EOF {
		WriteJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Request body must contain a single JSON object",
		})
		return
	}

	if strings.TrimSpace(newBook.Title) == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Title is required",
		})
		return
	}

	if strings.TrimSpace(newBook.Author) == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Author is required",
		})
		return
	}

	newBook.Title = strings.TrimSpace(newBook.Title)
	newBook.Author = strings.TrimSpace(newBook.Author)

	createdBook := service.CreateBook(newBook)

	WriteJSON(w, http.StatusCreated, createdBook)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := service.GetBooks()

	WriteJSON(w, http.StatusOK, books)
}
