package handler

import "net/http"

type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func Health(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status:  "ok",
		Message: "Library System API is healthy",
	}

	WriteJSON(w, http.StatusOK, response)
}
