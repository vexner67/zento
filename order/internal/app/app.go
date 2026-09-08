package app

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	orderv1 "github.com/vexner67/zento/order/api/order/v1"
	"github.com/vexner67/zento/order/internal/config"
	ordergrpc "github.com/vexner67/zento/order/internal/grpc"
	"github.com/vexner67/zento/order/internal/logger"
	"github.com/vexner67/zento/order/internal/postgres"
	"github.com/vexner67/zento/order/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type App struct {
	logger *slog.Logger
	pool   *pgxpool.Pool
	server *grpc.Server
	health *health.Server
	addr   string
}

func New(ctx context.Context, cfg config.Config) (*App, error) {
	logger, err := logger.New(cfg)
	if err != nil {
		return nil, fmt.Errorf("create logger: %w", err)
	}

	databaseCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	pool, err := postgres.New(databaseCtx, cfg.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("connect to postgres: %w", err)
	}

	repo := postgres.NewRepository(pool)
	svc := service.NewService(repo)
	handler := ordergrpc.NewHandler(svc, logger)

	server := grpc.NewServer()
	orderv1.RegisterOrderServiceServer(server, handler)

	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(server, healthServer)
	reflection.Register(server)

	return &App{
		logger: logger,
		pool:   pool,
		server: server,
		health: healthServer,
		addr:   fmt.Sprintf(":%d", cfg.GRPCPort),
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	listener, err := net.Listen("tcp", a.addr)
	if err != nil {
		return fmt.Errorf("listen on %s: %w", a.addr, err)
	}
	defer func() {
		if err = listener.Close(); err != nil {
			a.logger.Error("listen close", "error", err)
		}
	}()

	a.health.SetServingStatus(
		"",
		grpc_health_v1.HealthCheckResponse_SERVING,
	)

	serveErr := make(chan error, 1)
	go func() {
		serveErr <- a.server.Serve(listener)
	}()

	a.logger.Info("gRPC server started", "address", a.addr)

	select {
	case err := <-serveErr:
		if err != nil {
			return fmt.Errorf("serve gRPC: %w", err)
		}
		return nil

	case <-ctx.Done():
		a.logger.Info("stopping gRPC server")
	}

	a.health.Shutdown()

	timer := time.AfterFunc(5*time.Second, a.server.Stop)
	defer timer.Stop()

	a.server.GracefulStop()
	<-serveErr

	a.logger.Info("gRPC server stopped")
	return nil
}

func (a *App) Close() {
	a.pool.Close()
}
