package logrus

import (
	"io"
	"os"

	"github.com/palano/logger"
	"github.com/sirupsen/logrus"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

type logrusLogger struct {
	Logger *logrus.Logger
}

func (l *logrusLogger) Debug(args ...interface{}) {
	l.Logger.Debug(args...)
}

func (l *logrusLogger) Debugf(format string, args ...interface{}) {
	l.Logger.Debugf(format, args...)
}

func (l *logrusLogger) Info(args ...interface{}) {
	l.Logger.Info(args...)
}

func (l *logrusLogger) Infof(format string, args ...interface{}) {
	l.Logger.Infof(format, args...)
}

func (l *logrusLogger) Warn(args ...interface{}) {
	l.Logger.Warn(args...)
}

func (l *logrusLogger) Warnf(format string, args ...interface{}) {
	l.Logger.Warnf(format, args...)
}

func (l *logrusLogger) Error(args ...interface{}) {
	l.Logger.Error(args...)
}

func (l *logrusLogger) Errorf(format string, args ...interface{}) {
	l.Logger.Errorf(format, args...)
}

func (l *logrusLogger) Panic(args ...interface{}) {
	l.Logger.Panic(args...)
}

func (l *logrusLogger) Panicf(format string, args ...interface{}) {
	l.Logger.Panicf(format, args...)
}

func (l *logrusLogger) Fatal(args ...interface{}) {
	l.Logger.Fatal(args...)
}

func (l *logrusLogger) Fatalf(format string, args ...interface{}) {
	l.Logger.Fatalf(format, args...)
}

func newProvider(logger *logrus.Logger) (logger.Logger, error) {
	return &logrusLogger{
		Logger: logger,
	}, nil
}

func New(lc logger.LogConfig) logger.Logger {
	lcc := logger.CleanConfig(lc)
	logrusLog := logrus.New()

	if lcc.EnableFileLog {
		lbj := &lumberjack.Logger{
			Filename:   lcc.FileLog,
			MaxBackups: lcc.MaxBackups,
			MaxSize:    lcc.MaxSize,
			MaxAge:     lcc.MaxAge,
			Compress:   true,
		}
		writer := io.MultiWriter(os.Stderr, lbj)
		logrusLog.SetOutput(writer)
	}

	logrusLog.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyLevel: lcc.LevelKey,
			logrus.FieldKeyTime:  lcc.TimeKey,
			logrus.FieldKeyMsg:   lcc.MessageKey,
		},
	})

	logrusLog.AddHook(NewHook(lcc))

	log, _ := newProvider(logrusLog)
	logger.ReplaceProvider(log)

	return logger.NewLogger()
}
