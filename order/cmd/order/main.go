package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5"
	"github.com/vexner67/zento/order/internal/config"
	"github.com/vexner67/zento/order/internal/database"
	"github.com/vexner67/zento/order/internal/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config:", err)
		os.Exit(1)
	}

	logger, err := logger.New(cfg)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to create logger:", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	db, err := database.NewPostgres(ctx, cfg.DatabaseURL)
	cancel()

	if err != nil {
		logger.Error("failed to connect to postgres", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	addr := fmt.Sprintf(":%d", cfg.GRPCPort)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Error("failed to listen", "error", err)
		return
	}

	grpcServer := grpc.NewServer()

	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)

	reflection.Register(grpcServer)

	healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)

	logger.Info("starting gRPC server", "address", addr)

	if err = grpcServer.Serve(listener); err != nil {
		logger.Error("gRPC server stopped", "error", err)
	}
}
