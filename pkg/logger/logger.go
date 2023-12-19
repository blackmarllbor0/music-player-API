package logger

import "os"

type Logger interface {
	SetupLogger(logDir string) (*os.File, error)
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
}
