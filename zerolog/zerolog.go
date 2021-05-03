package zerolog

import (
	"io"
	"os"
	"time"

	"github.com/palano/logger"
	"github.com/rs/zerolog"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

type zerologLogger struct {
	logger *zerolog.Logger
}

func (l *zerologLogger) Debug(args ...interface{}) {
	log := l.logger.Debug()
	writeLog(log, args...)
}

func (l *zerologLogger) Debugf(format string, args ...interface{}) {
	log := l.logger.Debug()
	writeLog(log, args...)
}

func (l *zerologLogger) Info(args ...interface{}) {
	log := l.logger.Info()
	writeLog(log, args...)
}

func (l *zerologLogger) Infof(format string, args ...interface{}) {
	log := l.logger.Info()
	writeLog(log, args...)
}

func (l *zerologLogger) Warn(args ...interface{}) {
	log := l.logger.Warn()
	writeLog(log, args...)
}

func (l *zerologLogger) Warnf(format string, args ...interface{}) {
	log := l.logger.Warn()
	writeLog(log, args...)
}

func (l *zerologLogger) Error(args ...interface{}) {
	log := l.logger.Error()
	writeLog(log, args...)
}

func (l *zerologLogger) Errorf(format string, args ...interface{}) {
	log := l.logger.Error()
	writeLog(log, args...)
}

func (l *zerologLogger) Panic(args ...interface{}) {
	log := l.logger.Panic()
	writeLog(log, args...)
}

func (l *zerologLogger) Panicf(format string, args ...interface{}) {
	log := l.logger.Panic()
	writeLog(log, args...)
}

func (l *zerologLogger) Fatal(args ...interface{}) {
	log := l.logger.Fatal()
	writeLog(log, args...)
}

func (l *zerologLogger) Fatalf(format string, args ...interface{}) {
	log := l.logger.Fatal()
	writeLog(log, args...)
}

func (l *zerologLogger) WithFields(fields logger.Fields) logger.Logger {
	newLogger := l.logger.With().Fields(fields).Logger()
	return &zerologLogger{&newLogger}
}

func writeLog(log *zerolog.Event, args ...interface{}) {
	if len(args) > 0 && args[0] != nil {
		log.Msgf("%v", args...)
	} else {
		log.Msg("-")
	}
}

func newAdapter(logger *zerolog.Logger) (logger.Logger, error) {
	return &zerologLogger{
		logger: logger,
	}, nil
}

func New(lc logger.LogConfig) logger.Logger {
	lcc := logger.CleanConfig(lc)

	zerolog.LevelFieldName = lcc.LevelKey
	zerolog.TimestampFieldName = lcc.TimeKey
	zerolog.MessageFieldName = lcc.MessageKey
	zerolog.TimeFieldFormat = time.RFC3339

	writer := io.MultiWriter(os.Stderr)

	if lcc.EnableFileLog {
		lbj := &lumberjack.Logger{
			Filename:   lcc.FileLog,
			MaxBackups: lcc.MaxBackups,
			MaxSize:    lcc.MaxSize,
			MaxAge:     lcc.MaxAge,
			Compress:   true,
		}
		writer = io.MultiWriter(os.Stderr, lbj)
	}

	zeroLog := zerolog.New(writer).With().Timestamp().Logger().Hook(ZerologHook{
		LogConfig: lcc,
		CallerHook: []zerolog.Level{
			zerolog.DebugLevel,
			zerolog.InfoLevel,
			zerolog.WarnLevel,
			zerolog.ErrorLevel,
			zerolog.FatalLevel,
			zerolog.PanicLevel,
			zerolog.NoLevel,
		},
		StackLevels: []zerolog.Level{
			zerolog.PanicLevel,
			zerolog.FatalLevel,
			zerolog.ErrorLevel,
		},
	})

	log, _ := newAdapter(&zeroLog)
	logger.ReplaceProvider(log)

	return logger.NewLogger()
}
