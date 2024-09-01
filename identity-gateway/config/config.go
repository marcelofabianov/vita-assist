package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Project string
	Name    string
	ID      string
	VERSION string
	ENV     string
	TZ      string
}

func NewConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	cfg := &Config{
		Project: os.Getenv("PROJECT"),
		Name:    os.Getenv("NAME"),
		ID:      os.Getenv("ID"),
		VERSION: os.Getenv("VERSION"),
		ENV:     os.Getenv("ENV"),
		TZ:      os.Getenv("TZ"),
	}

	return cfg, nil
}
