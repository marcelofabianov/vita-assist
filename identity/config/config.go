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
	Log     Log
}

type Log struct {
	Level    string
	Format   string
	Output   string
	FilePath string
}

func NewConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	cfg := &Config{
		Project: os.Getenv("IDE_PROJECT"),
		Name:    os.Getenv("IDE_NAME"),
		ID:      os.Getenv("IDE_ID"),
		VERSION: os.Getenv("IDE_VERSION"),
		ENV:     os.Getenv("IDE_ENV"),
		TZ:      os.Getenv("IDE_TZ"),
		Log: Log{
			Level:    os.Getenv("IDE_LOG_LEVEL"),
			Format:   os.Getenv("IDE_LOG_FORMAT"),
			Output:   os.Getenv("IDE_LOG_OUTPUT"),
			FilePath: os.Getenv("IDE_LOG_FILE_PATH"),
		},
	}

	return cfg, nil
}
