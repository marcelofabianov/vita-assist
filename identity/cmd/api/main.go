package main

import (
	"log/slog"

	"github.com/marcelofabianov/vita-assist/identity/config"
	"github.com/marcelofabianov/vita-assist/identity/pkg/logger"
)

type App struct {
	config *config.Config
	logger *logger.Logger
}

func main() {
	app, err := build()

	if err != nil {
		slog.Error("Failed to start the identity",
			slog.String("error", err.Error()),
			slog.String("context", "main function"),
		)
		return
	}
	defer app.logger.Close()

	app.logger.Info("Starting the identity",
		app.logger.String("status", "running"),
		app.logger.String("env", app.config.ENV),
	)
}

func build() (*App, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	slog.Info("Configuration loaded",
		slog.String("ID", cfg.ID),
		slog.String("ENV", cfg.ENV),
	)

	logger, err := logger.NewLogger(cfg.Log)
	if err != nil {
		slog.Error("Failed to create logger",
			slog.String("error", err.Error()),
			slog.String("context", "main function"),
		)
		return nil, err
	}

	return &App{config: cfg, logger: logger}, nil
}
