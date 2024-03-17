package grpcApp

import (
	"fmt"
	"github.com/xjl0/mi-sso/internal/app/validation"
	authgRPC "github.com/xjl0/mi-sso/internal/grpc/auth"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	gRPCHost   string
	vl         *validation.App
}

func NewApp(log *slog.Logger, gRPCHost string, vl *validation.App) *App {
	gRPCServer := grpc.NewServer()

	authgRPC.Register(gRPCServer, vl)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		gRPCHost:   gRPCHost,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const o = "grpcApp.Run"
	log := a.log.With(slog.String("o", o))

	log.Info("GRPC server starting")

	l, err := net.Listen("tcp", a.gRPCHost)
	if err != nil {
		return fmt.Errorf("%s: %w", o, err)
	}

	log.Info("GRPC server started", slog.String("address", l.Addr().String()))
	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", o, err)
	}

	return nil
}

func (a *App) Stop() {
	const o = "grpcApp.Stop"
	log := a.log.With(slog.String("o", o))
	log.Info("GRPC server stopping")
	a.gRPCServer.GracefulStop()
}
