package httpserver

import (
	"net/http"

	"library-system/internal/handler"
)

func NewRouter(bookHandler *handler.BookHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handler.Health)

	mux.HandleFunc("/api/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			bookHandler.GetBooks(w, r)

		case http.MethodPost:
			bookHandler.CreateBook(w, r)

		default:
			handler.WriteJSON(w, http.StatusMethodNotAllowed, map[string]string{
				"error": "Method not allowed",
			})
		}
	})

	return loggingMiddleware(
		corsMiddleware(mux),
	)
}
