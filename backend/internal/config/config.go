package config

import "os"

type Config struct {
	AppEnv  string
	AppPort string
}

func Load() Config {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "development"
	}

	appPort := os.Getenv("APP_PORT")
	if appEnv == "" {
		appEnv = "8081"
	}

	return Config{
		AppEnv:  appEnv,
		AppPort: appPort,
	}
}
