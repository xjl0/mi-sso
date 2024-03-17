package app

import (
	grpcApp "github.com/xjl0/mi-sso/internal/app/grpc"
	"github.com/xjl0/mi-sso/internal/app/validation"
	"log/slog"
	"time"
)

type App struct {
	GRPCServer *grpcApp.App
}

func NewApp(log *slog.Logger, GRPCHost string, tokenTTL time.Duration) *App {
	vl := validation.New()
	gApp := grpcApp.NewApp(log, GRPCHost, vl)
	return &App{
		GRPCServer: gApp,
	}
}
