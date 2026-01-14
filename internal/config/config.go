package config

import (
	"os"

	pc "github.com/nightmaker00/go-tasks-api/pkg/db/postgres"
)

type Config struct {
	Server struct {
		Address string
		Port    string
	}
	pc.Config
}

func Load() (*Config, error) {
	cfg := &Config{}

	//default address
	cfg.Server.Address = "0.0.0.0"
	cfg.Server.Port = "8080"

	cfg.Config.Host = "localhost"
	cfg.Config.Port = "5432"
	cfg.Config.User = "postgres"
	cfg.Config.Password = "postgres"
	cfg.Config.DBName = "tasks"
	cfg.Config.SSLMode = "disable"

	//in env if exists
	if host := os.Getenv("SERVER_HOST"); host != "" {
		cfg.Server.Address = host
	}

	if port := os.Getenv("SERVER_PORT"); port != "" {
		cfg.Server.Port = port
	}

	if host := os.Getenv("POSTGRES_HOST"); host != "" {
		cfg.Config.Host = host
	}
	if port := os.Getenv("POSTGRES_PORT"); port != "" {
		cfg.Config.Port = port
	}
	if user := os.Getenv("POSTGRES_USER"); user != "" {
		cfg.Config.User = user
	}
	if password := os.Getenv("POSTGRES_PASSWORD"); password != "" {
		cfg.Config.Password = password
	}
	if name := os.Getenv("POSTGRES_DB"); name != "" {
		cfg.Config.DBName = name
	}
	if ssl := os.Getenv("POSTGRES_SSLMODE"); ssl != "" {
		cfg.Config.SSLMode = ssl
	}

	return cfg, nil
}
