package utils

import (
	"context"

	"github.com/Fiagram/account_service/internal/configs"
	"go.uber.org/zap"
)

func getZapLoggerLevel(level string) zap.AtomicLevel {
	switch level {
	case "debug":
		return zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		return zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		return zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "panic":
		return zap.NewAtomicLevelAt(zap.PanicLevel)
	default:
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	}
}

func InitializeLogger(logConfig configs.Log) (logger *zap.Logger, cleanup func(), err error) {
	zapLoggerConfig := zap.NewProductionConfig()
	zapLoggerConfig.Level = getZapLoggerLevel(logConfig.Level)

	logger, err = zapLoggerConfig.Build()
	if err != nil {
		return nil, nil, err
	}

	cleanup = func() {
		// delibrately ignore the returned error here
		_ = logger.Sync()
	}

	return
}

func LoggerWithContext(_ context.Context, logger *zap.Logger) *zap.Logger {
	// TODO: Add request ID to context
	return logger
}
