package main

import (
	"log"
	"net/http"

	"library-system/internal/config"
	"library-system/internal/httpserver"
)

func main() {
	cfg := config.Load()

	router := httpserver.NewRouter()

	server := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: router,
	}

	log.Printf(
		"Server running in %s mode on http://localhost:%s",
		cfg.AppEnv,
		cfg.AppPort,
	)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
