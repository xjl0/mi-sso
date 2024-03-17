package main

import (
	"log/slog"
	"mi-sso/internal/config"
	"os"
)

func main() {
	cfg := config.MustLoad()
	log := setLogger(cfg.LogLevel, cfg.IsLocal)

	log.Info("Starting...", slog.String("creator", "github.com/xjl0"))

}

func setLogger(level string, isLocal bool) *slog.Logger {
	var lvl slog.Leveler
	switch level {
	case "debug":
		lvl = slog.LevelDebug
	case "info":
		lvl = slog.LevelInfo
	case "warn":
		lvl = slog.LevelWarn
	default:
		lvl = slog.LevelError
	}

	if isLocal {
		return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: lvl}))
	}

	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: lvl}))
}
