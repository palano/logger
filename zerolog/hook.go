package zerolog

import (
	"github.com/palano/logger"
	"github.com/palano/stacktracer"
	"github.com/rs/zerolog"
)

type ZerologHook struct {
	CallerHook  []zerolog.Level
	StackLevels []zerolog.Level
	LogConfig   logger.LogConfig
}

func (h ZerologHook) Run(e *zerolog.Event, l zerolog.Level, msg string) {
	if h.LogConfig.EnableCaller {
		stack := stacktracer.Caller(5)
		for _, level := range h.CallerHook {
			if level == l {
				e.Str(h.LogConfig.CallerKey, stack.String())
			}
		}
	}

	if h.LogConfig.EnableStacktracer {
		stacks := stacktracer.Callers(5)
		for _, level := range h.CallerHook {
			if level == l {
				e.Str(h.LogConfig.StacktracerKey, stacks.String())
			}
		}
	}
}
