package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"library-system/internal/config"
	"library-system/internal/database"
	"library-system/internal/handler"
	"library-system/internal/httpserver"
	"library-system/internal/repository"
	"library-system/internal/service"
)

func main() {
	cfg := config.Load()

	var bookRepository service.BookRepository

	if cfg.DBPassword == "" {
		log.Println("No database password configured. Using in-memory repository.")

		bookRepository = repository.NewMemoryBookRepository()
	} else {
		db, err := database.Connect(database.Config{
			Host:     cfg.DBHost,
			Port:     cfg.DBPort,
			User:     cfg.DBUser,
			Password: cfg.DBPassword,
			Name:     cfg.DBName,
		})
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}
		defer db.Close()

		log.Println("Connected to SQL Server successfully")

		bookRepository = repository.NewBookRepository(db)
	}

	bookService := service.NewBookService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	router := httpserver.NewRouter(bookHandler)

	server := &http.Server{
		Addr:         ":" + cfg.AppPort,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Printf(
			"Server running in %s mode on http://localhost:%s",
			cfg.AppEnv,
			cfg.AppPort,
		)

		if err := server.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)

	signal.Notify(
		stop,
		os.Interrupt,
		syscall.SIGTERM,
	)

	<-stop

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown failed: %v", err)
	}

	log.Println("Server stopped")
}
