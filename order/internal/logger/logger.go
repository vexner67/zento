package logger

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/vexner67/zento/order/internal/config"
)

func New(cfg config.Config) (*slog.Logger, error) {
	level, err := parseLevel(cfg.LogLevel)
	if err != nil {
		return nil, err
	}

	opts := &slog.HandlerOptions{
		Level:     level,
		AddSource: true,
	}

	var handler slog.Handler

	switch cfg.LogFormat {
	case "text":
		handler = slog.NewTextHandler(os.Stdout, opts)
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, opts)
	default:
		return nil, fmt.Errorf("unsupported log format: %q", cfg.LogFormat)
	}

	return slog.New(handler), nil
}

func parseLevel(level string) (slog.Level, error) {
	switch level {
	case "debug":
		return slog.LevelDebug, nil
	case "info":
		return slog.LevelInfo, nil
	case "warn":
		return slog.LevelWarn, nil
	case "error":
		return slog.LevelError, nil
	default:
		return 0, fmt.Errorf("unsupported log level: %q", level)
	}
}
