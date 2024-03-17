package grpcApp

import (
	"fmt"
	authgRPC "github.com/xjl0/mi-sso/internal/grpc/auth"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func NewApp(log *slog.Logger, port int) *App {
	gRPCServer := grpc.NewServer()

	authgRPC.Register(gRPCServer)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const o = "grpcApp.Run"
	log := a.log.With(slog.String("o", o), slog.Int("port", a.port))

	log.Info("Starting gRPC server...")

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", o, err)
	}

	log.Info("gRPC server started", slog.String("address", l.Addr().String()))
	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", o, err)
	}

	return nil
}

func (a *App) Stop() {
	const o = "grpcApp.Stop"
	log := a.log.With(slog.String("o", o))
	log.Info("Stopping gRPC server...")
	a.gRPCServer.GracefulStop()
}
