package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	GRPCPort  int    `env:"GRPC_PORT" required:"true" default:"50051"`
	LogLevel  string `env:"LOG_LEVEL" required:"true" default:"info"`
	LogFormat string `env:"LOG_FORMAT" required:"true" default:"text"`
}

func Load() (Config, error) {
	if err := godotenv.Load(); err != nil {
		return Config{}, fmt.Errorf("load .env: %w", err)
	}

	cfg, err := env.ParseAs[Config]()
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
