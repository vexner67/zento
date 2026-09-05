package main

import (
	"log"

	"github.com/vexner67/zento/order/internal/config"
	"github.com/vexner67/zento/order/internal/logger"
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

	logger.Info("starting order service")
}
