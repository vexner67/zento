package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	GRPCPort  int    `env:"GRPC_PORT,required"`
	LogLevel  string `env:"LOG_LEVEL,required"`
	LogFormat string `env:"LOG_FORMAT,required"`

	DatabaseURL string `env:"DATABASE_URL,required"`
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
