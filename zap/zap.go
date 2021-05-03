package zap

import (
	"fmt"
	"net/url"

	"github.com/palano/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

type lumberjackSink struct {
	*lumberjack.Logger
}

func (lumberjackSink) Sync() error { return nil }

type zapLogger struct {
	sugaredLogger *zap.SugaredLogger
}

func (l *zapLogger) Debug(args ...interface{}) {
	l.sugaredLogger.Debug(args...)
}

func (l *zapLogger) Debugf(format string, args ...interface{}) {
	l.sugaredLogger.Debugf(format, args...)
}

func (l *zapLogger) Info(args ...interface{}) {
	l.sugaredLogger.Info(args...)
}

func (l *zapLogger) Infof(format string, args ...interface{}) {
	l.sugaredLogger.Infof(format, args...)
}

func (l *zapLogger) Warn(args ...interface{}) {
	l.sugaredLogger.Warn(args...)
}

func (l *zapLogger) Warnf(format string, args ...interface{}) {
	l.sugaredLogger.Warnf(format, args...)
}

func (l *zapLogger) Error(args ...interface{}) {
	l.sugaredLogger.Error(args...)
}

func (l *zapLogger) Errorf(format string, args ...interface{}) {
	l.sugaredLogger.Errorf(format, args...)
}

func (l *zapLogger) Panic(args ...interface{}) {
	l.sugaredLogger.Panic(args...)
}

func (l *zapLogger) Panicf(format string, args ...interface{}) {
	l.sugaredLogger.Panicf(format, args...)
}

func (l *zapLogger) Fatal(args ...interface{}) {
	l.sugaredLogger.Fatal(args...)
}

func (l *zapLogger) Fatalf(format string, args ...interface{}) {
	l.sugaredLogger.Fatalf(format, args...)
}

func (l *zapLogger) WithFields(fields logger.Fields) logger.Logger {
	var f = make([]interface{}, 0)
	for k, v := range fields {
		f = append(f, k)
		f = append(f, v)
	}
	newLogger := l.sugaredLogger.With(f...)
	return &zapLogger{newLogger}
}

func newProvider(logger *zap.Logger) (logger.Logger, error) {
	sugaredLogger := logger.WithOptions(zap.AddCallerSkip(1)).Sugar()
	return &zapLogger{
		sugaredLogger: sugaredLogger,
	}, nil
}

func New(cl logger.LogConfig) logger.Logger {
	lcc := logger.CleanConfig(cl)
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.LevelKey = lcc.LevelKey
	cfg.EncoderConfig.TimeKey = lcc.TimeKey
	cfg.EncoderConfig.MessageKey = lcc.MessageKey
	cfg.EncoderConfig.CallerKey = lcc.CallerKey
	cfg.EncoderConfig.StacktraceKey = lcc.StacktracerKey
	cfg.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	cfg.DisableCaller = !lcc.EnableCaller
	cfg.DisableStacktrace = !lcc.EnableStacktracer

	// Check if file log enable
	if lcc.EnableFileLog {
		zap.RegisterSink("lumberjack", func(u *url.URL) (zap.Sink, error) {
			return lumberjackSink{
				Logger: &lumberjack.Logger{
					Filename:   lcc.FileLog,
					MaxBackups: lcc.MaxBackups,
					MaxSize:    lcc.MaxSize,
					MaxAge:     lcc.MaxAge,
					Compress:   true,
				},
			}, nil
		})
		cfg.OutputPaths = append(cfg.OutputPaths, fmt.Sprintf("lumberjack:%s", lcc.FileLog))
	}

	zapLog, _ := cfg.Build()
	log, _ := newProvider(zapLog)
	logger.ReplaceProvider(log)

	return logger.NewLogger()
}
