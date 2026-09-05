package main

import (
	"fmt"
	"log"
	"net"

	"github.com/vexner67/zento/order/internal/config"
	"github.com/vexner67/zento/order/internal/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	logger, err := logger.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	addr := fmt.Sprintf(":%d", cfg.GRPCPort)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Error("failed to listen", "error", err)
		return
	}

	grpcServer := grpc.NewServer()

	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)

	healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)

	logger.Info("starting gRPC server", "address", addr)

	if err = grpcServer.Serve(listener); err != nil {
		logger.Error("gRPC server stopped", "error", err)
	}
}
