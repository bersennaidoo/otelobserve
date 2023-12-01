package zlog

import (
	"go.uber.org/zap"
)

func NewZlog(name string) *zap.Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger = logger.Named(name)

	return logger
}
