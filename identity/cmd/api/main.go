package main

import (
	"log/slog"

	"github.com/marcelofabianov/vita-assist/identity/config"
)

func main() {
	cfg, err := run()

	if err != nil {
		slog.Error("Failed to start the identity gateway",
			slog.String("error", err.Error()),
			slog.String("context", "main function"),
		)
		return
	}

	slog.Info("Identity gateway is running",
		slog.String("ID", cfg.ID),
		slog.String("status", "running"),
	)
}

func run() (*config.Config, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	slog.Info("Configuration loaded",
		slog.String("ID", cfg.ID),
		slog.String("ENV", cfg.ENV),
	)

	return cfg, nil
}
