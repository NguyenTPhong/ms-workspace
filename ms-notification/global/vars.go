package global

import (
	"go.uber.org/zap"
	"ms-workspace/ms-notification/global/config"
	"ms-workspace/package/logger"
)

var (
	Logger *zap.Logger
)

func Init() {
	Logger = logger.NewLogger(config.LogLevel)
}

func DeInit() {
	Logger.Sync()
}
