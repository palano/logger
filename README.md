# logger

The modular log for log providers

## Installation

`go get -u github.com/palano/logger`

## Quick Start

```go
// LogConfig
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
```

### Example for :zap: Zap

```go

import (
    "github.com/palano/logger"
    "github.com/palano/logger/zap"
)

func main() {
    lc := logger.LogConfig{}
    log := zap.New(lc)
    log.Info("Example")
    
    log.WithFields(logger.Fields{
        "content": "example",
    }).Info("Example")
}
```

### Example for Logrus

```go

import (
    "github.com/palano/logger"
    "github.com/palano/logger/logrus"
)

func main() {
    lc := logger.LogConfig{}
    log := logrus.New(lc)
    log.Info("Example")
    
    log.WithFields(logger.Fields{
        "content": "example",
    }).Info("Example")
}
```

### Example for Zerolog

```go

import (
    "github.com/palano/logger"
    "github.com/palano/logger/zerolog"
)

func main() {
    lc := logger.LogConfig{}
    log := zerolog.New(lc)
    log.Info("Example")
    
    log.WithFields(logger.Fields{
        "content": "example",
    }).Info("Example")
}
```

## Providers

- :zap: [uber-go/zap](https://github.com/uber-go/zap)
- [sirupsen/logrus](https://github.com/sirupsen/logrus)
- [rs/zerolog](https://github.com/rs/zerolog)

## File Log Writer
- os.Stderr
- [lumberjack](gopkg.in/natefinch/lumberjack.v2)
## License
Released under the [MIT License](LICENSE).