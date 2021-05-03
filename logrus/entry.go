package logrus

import (
	"github.com/sirupsen/logrus"
)

type logrusEntry struct {
	Entry *logrus.Entry
}

func (l *logrusEntry) Debug(args ...interface{}) {
	l.Entry.Debug(args...)
}

func (l *logrusEntry) Debugf(format string, args ...interface{}) {
	l.Entry.Debugf(format, args...)
}

func (l *logrusEntry) Info(args ...interface{}) {
	l.Entry.Info(args...)
}

func (l *logrusEntry) Infof(format string, args ...interface{}) {
	l.Entry.Infof(format, args...)
}

func (l *logrusEntry) Warn(args ...interface{}) {
	l.Entry.Warn(args...)
}

func (l *logrusEntry) Warnf(format string, args ...interface{}) {
	l.Entry.Warnf(format, args...)
}

func (l *logrusEntry) Error(args ...interface{}) {
	l.Entry.Error(args...)
}

func (l *logrusEntry) Errorf(format string, args ...interface{}) {
	l.Entry.Errorf(format, args...)
}

func (l *logrusEntry) Panic(args ...interface{}) {
	l.Entry.Panic(args...)
}

func (l *logrusEntry) Panicf(format string, args ...interface{}) {
	l.Entry.Panicf(format, args...)
}

func (l *logrusEntry) Fatal(args ...interface{}) {
	l.Entry.Fatal(args...)
}

func (l *logrusEntry) Fatalf(format string, args ...interface{}) {
	l.Entry.Fatalf(format, args...)
}
