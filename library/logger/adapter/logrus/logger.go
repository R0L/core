package logrus

import (
	"github.com/R0L/core/library/logger/adapter/base"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type Logger struct {
	base.Log
	*log.Logger
}

func NewLogger(cf *base.Config) *Logger {
	var l *Logger
	l = &Logger{}
	l.Config = cf
	return l
}

// Init
func (l *Logger) Init() error {
	l.Logger = log.New()
	l.getLogger().SetOutput(l.Config.Out)
	if err := l.setLevel(l.Config.Level); nil != err {
		return errors.Wrap(err, "logrus Init err")
	}
	l.setFormatter(l.Config.Format)
	l.getLogger().SetReportCaller(true)
	return nil
}

// getLogger
func (l *Logger) getLogger() *log.Logger {
	return l.Logger
}

// setLevel
func (l *Logger) setLevel(logLevel base.LogLevel) error {
	level, err := log.ParseLevel(string(logLevel))
	if nil != err {
		return errors.Wrap(err, "logrus setLevel error")
	}
	l.getLogger().SetLevel(level)
	return nil
}

// setFormatter
func (l *Logger) setFormatter(format base.LogFormat) {
	var formatter log.Formatter
	switch format {
	case base.JSON:
		formatter = &log.JSONFormatter{}
	case base.TEXT:
		formatter = &log.TextFormatter{}
	}
	l.getLogger().SetFormatter(formatter)
}

func (l *Logger) Info(args ...interface{}) {
	l.getLogger().Info(args...)
}
func (l *Logger) Infof(format string, args ...interface{}) {
	l.getLogger().Infof(format, args...)
}
func (l *Logger) Infoln(args ...interface{}) {
	l.getLogger().Infoln(args...)
}
func (l *Logger) Debug(args ...interface{}) {
	l.getLogger().Debug(args...)
}
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.getLogger().Debugf(format, args...)
}
func (l *Logger) Debugln(args ...interface{}) {
	l.getLogger().Debugln(args...)
}
func (l *Logger) Error(args ...interface{}) {
	l.getLogger().Error(args...)
}
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.getLogger().Errorf(format, args...)
}
func (l *Logger) Errorln(args ...interface{}) {
	l.getLogger().Errorln(args...)
}
