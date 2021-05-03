package logrus

import (
	"github.com/palano/logger"
	"github.com/palano/stacktracer"
	"github.com/sirupsen/logrus"
)

type LogrusHook struct {
	CallerHook  []logrus.Level
	StackLevels []logrus.Level
	LogConfig   logger.LogConfig
}

func NewHook(lc logger.LogConfig) LogrusHook {
	return LogrusHook{
		CallerHook:  logrus.AllLevels,
		StackLevels: []logrus.Level{logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel},
		LogConfig:   lc,
	}
}

func (h LogrusHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h LogrusHook) Fire(entry *logrus.Entry) error {
	var skip int
	if len(entry.Data) == 0 {
		skip = 8
	} else {
		skip = 6
	}

	if h.LogConfig.EnableCaller {
		for _, level := range h.CallerHook {
			stack := stacktracer.Caller(skip)
			if entry.Level == level {
				entry.Data[h.LogConfig.CallerKey] = stack.String()
				break
			}
		}
	}

	if h.LogConfig.EnableStacktracer {
		frames := stacktracer.Callers(skip)
		for _, level := range h.StackLevels {
			if entry.Level == level && h.LogConfig.EnableStacktracer {
				entry.Data[h.LogConfig.StacktracerKey] = frames.String()
				break
			}
		}
	}

	return nil
}
