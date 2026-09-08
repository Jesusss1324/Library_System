package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/microsoft/go-mssqldb"
)

// Config contains everything required to connect to SQL Server.
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

// Connect creates and verifies a connection to SQL Server.
func Connect(cfg Config) (*sql.DB, error) {
	connectionString := fmt.Sprintf(
		"sqlserver://%s:%s@%s:%s?database=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)

	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}
