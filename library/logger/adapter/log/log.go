package log

import (
	"github.com/R0L/core/library/logger/adapter/base"
	"log"
)

type Log struct {
	base.Log
	*log.Logger
}

func NewLog(cf *base.Config) *Log {
	var l = &Log{}
	l.Config = cf
	return l
}

/**
支持 config.Out
*/

func (l *Log) Init() error {
	l.Logger = log.New(l.Out, "", log.LstdFlags|log.Lshortfile)
	return nil
}

func (l *Log) Info(args ...interface{}) {
	l.Logger.Print(args...)
}
func (l *Log) Infof(format string, args ...interface{}) {
	l.Logger.Printf(format, args...)
}
func (l *Log) Infoln(args ...interface{}) {
	l.Logger.Println(args...)
}
func (l *Log) Debug(args ...interface{}) {
	l.Logger.Print(args...)
}
func (l *Log) Debugf(format string, args ...interface{}) {
	l.Logger.Printf(format, args...)
}
func (l *Log) Debugln(args ...interface{}) {
	l.Logger.Println(args...)
}
func (l *Log) Error(args ...interface{}) {
	l.Logger.Fatal(args...)
}
func (l *Log) Errorf(format string, args ...interface{}) {
	l.Logger.Fatalf(format, args...)
}
func (l *Log) Errorln(args ...interface{}) {
	l.Logger.Fatalln(args...)
}
