package config

import "os"

type Config struct {
	AppEnv     string
	AppPort    string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func Load() Config {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "development"
	}

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8081"
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "db"
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "1433"
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "sa"
	}

	dbPassword := os.Getenv("DB_PASSWORD")

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "library_db"
	}

	return Config{
		AppEnv:     appEnv,
		AppPort:    appPort,
		DBHost:     dbHost,
		DBPort:     dbPort,
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DBName:     dbName,
	}
}
