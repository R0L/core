package logger

import (
	"github.com/R0L/core/library/logger/adapter/base"
	"github.com/R0L/core/library/logger/adapter/log"
	"github.com/R0L/core/library/logger/adapter/logrus"
)

func GetAdapter(cfg *base.Config) Logger {
	var logger Logger
	switch cfg.Type {
	case base.LOGRUS:
		logger = logrus.NewLogger(cfg)
	default:
		logger = log.NewLog(cfg)
	}
	return logger
}
