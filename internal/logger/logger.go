package logger

import (
	"log/slog"
	"os"

	"github.com/talismanch1k/gorain/internal/closer"
)

func SetupLogger() (func(), error) {
	f, err := os.OpenFile("log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return nil, err
	}

	logger := slog.New(slog.NewTextHandler(f, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)

	fileClose := func() {
		slog.Info("Logger shutting down, file closing")
		closer.CloseOrLog(f)
	}

	return fileClose, nil
}
