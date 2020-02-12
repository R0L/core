package logger

import (
	"fmt"
	"github.com/R0L/core/library/file"
	logger "github.com/R0L/core/library/logger/adapter"
	"github.com/R0L/core/library/logger/adapter/base"
	"io"
	"os"
	"time"
)

var (
	log logger.Logger
)

func init() {

	f, err := file.OpenFile(config.Out, fmt.Sprintf("%s.log", time.Now().Format("2006-01-02")))
	if nil != err {
		log.Errorf("logger init CreateFile err, err: %v", err)
	}

	log = logger.GetAdapter(&base.Config{
		Type:   base.LOGRUS,
		Out:    io.MultiWriter(os.Stdout, &FileWrite{file: f}),
		Level:  base.INFO,
		Format: base.TEXT,
	})
	if err := log.Init(); nil != err {
		log.Errorf("logger init err, err: %v", err)
	}
}

func Info(args ...interface{}) {
	log.Info(args)
}
func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}
func Infoln(args ...interface{}) {
	log.Infoln(args...)
}
func Debug(args ...interface{}) {
	log.Debug(args...)
}
func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}
func Debugln(args ...interface{}) {
	log.Debugln(args...)
}
func Error(args ...interface{}) {
	log.Error(args...)
}
func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}
func Errorln(args ...interface{}) {
	log.Errorln(args...)
}
