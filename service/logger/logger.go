package logger

import "github.com/google/wire"

type Logger interface {
	Info(string)
	Error(string)
	Debug(string)
	Warn(string)
}

var defaultLogger Logger

func InitLogger() {
	defaultLogger = NewZapLog()
}

func NewLogger() Logger {
	return NewZapLog()
}

// SetLogger a customize logger to default logger
func SetLogger(logger Logger) {
	defaultLogger = logger
}

func Info(msg string) {
	defaultLogger.Info(msg)
}

func Error(msg string) {
	defaultLogger.Error(msg)
}

func Debug(msg string) {
	defaultLogger.Debug(msg)
}

func Warn(msg string) {
	defaultLogger.Info(msg)
}

var ProviderSet = wire.NewSet(NewLogger)
