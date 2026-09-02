package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("APP_PORT")

	if port == "" {
		port = "8081"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Library System API")
	})

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Server running on http://localhost:%s\n", port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
