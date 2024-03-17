package app

import (
	grpcApp "github.com/xjl0/mi-sso/internal/app/grpc"
	"log/slog"
	"time"
)

type App struct {
	GRPCServer *grpcApp.App
}

func NewApp(log *slog.Logger, GRPCHost string, tokenTTL time.Duration) *App {
	gApp := grpcApp.NewApp(log, GRPCHost)
	return &App{
		GRPCServer: gApp,
	}
}
