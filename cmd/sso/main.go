package main

import (
	"github.com/xjl0/mi-sso/internal/app"
	"github.com/xjl0/mi-sso/internal/config"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()
	log := setLogger(cfg.LogLevel, cfg.IsLocal)

	log.Info("App Starting", slog.String("creator", "github.com/xjl0"))

	apl := app.NewApp(log, cfg.GRPCHost, cfg.JwtTTL)

	go apl.GRPCServer.MustRun()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	si := <-stop

	log.Info("App Stopping", slog.String("signal", si.String()))

	apl.GRPCServer.Stop()

	log.Info("App Stopped")
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
