package base

import "io"

type LogType string
type LogLevel string
type LogFormat string

const (
	// logType
	DEFAULT LogType = ""
	LOGRUS  LogType = "logrus"

	// logLevel
	INFO  LogLevel = "info"
	DEBUG LogLevel = "debug"
	ERROR LogLevel = "error"

	// logFormat
	TEXT LogFormat = "text"
	JSON LogFormat = "json"
)

type Config struct {
	Type   LogType // LOGRUS, DEFAULT
	Out    io.Writer
	Level  LogLevel  // info，debug，error
	Format LogFormat // json，text
}

type Log struct {
	*Config
}
