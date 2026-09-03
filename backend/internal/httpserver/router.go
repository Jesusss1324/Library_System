package httpserver

import (
	"net/http"

	"library-system/internal/handler"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handler.Health)

	return loggingMiddleware(mux)
}
