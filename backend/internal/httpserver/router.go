package httpserver

import (
	"net/http"

	"library-system/internal/handler"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handler.Health)

	mux.HandleFunc("/api/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetBooks(w, r)

		case http.MethodPost:
			handler.CreateBook(w, r)

		default:
			handler.WriteJSON(w, http.StatusMethodNotAllowed, map[string]string{
				"error": "Method not allowed",
			})
		}
	})

	return mux
}
