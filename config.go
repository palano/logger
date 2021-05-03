package logger

import (
	"path"
)

// Log config
type LogConfig struct {
	LevelKey          string
	TimeKey           string
	MessageKey        string
	CallerKey         string
	StacktracerKey    string
	EnableCaller      bool
	EnableStacktracer bool
	EnableFileLog     bool
	FileLog           string
	MaxSize           int
	MaxBackups        int
	MaxAge            int
}

func CleanConfig(cfg LogConfig) LogConfig {
	if cfg.LevelKey == "" {
		cfg.LevelKey = "level"
	}

	if cfg.TimeKey == "" {
		cfg.TimeKey = "time"
	}

	if cfg.MessageKey == "" {
		cfg.MessageKey = "msg"
	}

	if cfg.CallerKey == "" {
		cfg.CallerKey = "caller"
	}

	if cfg.StacktracerKey == "" {
		cfg.StacktracerKey = "tracer"
	}

	if cfg.FileLog == "" {
		cfg.FileLog = path.Join("logs", "app.log")
	}

	return cfg
}
