package main

import (
	"log/slog"
	"os"

	"github.com/talismanch1k/gorain/internal/app"
	"github.com/talismanch1k/gorain/internal/logger"
)

func main() {
	fileClose, err := logger.SetupLogger()
	if err != nil {
		panic("Can't setup logger")
	}
	defer fileClose()

	slog.Info("Application started")

	if err = app.DrawScreen(); err != nil {
		slog.Error("Failed to draw screen", "err", err)
		os.Exit(1)
	}
}
