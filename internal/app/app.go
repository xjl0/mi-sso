package app

import (
	grpcApp "github.com/xjl0/mi-sso/internal/app/grpc"
	"log/slog"
	"time"
)

type App struct {
	GRPCServer *grpcApp.App
}

func NewApp(log *slog.Logger, grpcPort int, tokenTTL time.Duration) *App {
	gApp := grpcApp.NewApp(log, grpcPort)
	return &App{
		GRPCServer: gApp,
	}
}
