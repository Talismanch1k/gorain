package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/talismanch1k/gorain/internal/app"
	"github.com/talismanch1k/gorain/internal/logger"
)

func run() error {
	fileClose, err := logger.SetupLogger()
	if err != nil {
		return fmt.Errorf("initialize logger: %w", err)
	}
	defer fileClose()
	slog.Info("Application started")

	if err = app.DrawScreen(); err != nil {
		return fmt.Errorf("draw screen: %w", err)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		slog.Error("program failed", "err", err)
		os.Exit(1)
	}
}
