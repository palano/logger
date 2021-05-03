package logrus

import (
	"github.com/palano/logger"
	"github.com/sirupsen/logrus"
)

func (l *logrusLogger) WithFields(fields logger.Fields) logger.Logger {
	return &logrusEntry{
		Entry: l.Logger.WithFields(convertFields(fields)),
	}
}

func (l *logrusEntry) WithFields(fields logger.Fields) logger.Logger {
	return &logrusEntry{
		Entry: l.Entry.WithFields(convertFields(fields)),
	}
}

func convertFields(fields logger.Fields) logrus.Fields {
	logrusFields := logrus.Fields{}
	for index, val := range fields {
		logrusFields[index] = val
	}
	return logrusFields
}
