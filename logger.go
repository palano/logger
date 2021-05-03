package logger

import "sync"

type Fields map[string]interface{}

type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	WithFields(keyValues Fields) Logger
}

var (
	mu sync.RWMutex
	lo Logger
)

func ReplaceProvider(logger Logger) func() {
	mu.Lock()
	prev := lo
	lo = logger
	mu.Unlock()
	return func() { ReplaceProvider(prev) }
}

func NewLogger() Logger {
	mu.RLock()
	logger := lo
	mu.RUnlock()
	return logger
}
