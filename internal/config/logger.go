package config

import (
	"log/slog"
	"os"
	"strings"
)

func ConfigLogger() *slog.Logger {
	logLevel, ok := MapLog[strings.ToUpper(Configs.Logger.Level)]

	if !ok {
		logLevel = slog.LevelError
	}

	// TODO: check slog.HandlerOptions and to config.yaml
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		// AddSource: Configs.Logger.AddSrouce,
		Level: logLevel,
	}))

	return logger
}
