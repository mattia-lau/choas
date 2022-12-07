package logging

import (
	"chaos-go/internal/config"
	"log"

	"go.uber.org/zap"
)

func SetupLogger(cfg *config.Config) *zap.Logger {
	var (
		logger *zap.Logger
		err    error
	)

	if cfg.Env != config.Dev {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}

	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	defer logger.Sync()

	return logger
}
