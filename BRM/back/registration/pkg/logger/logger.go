package logger

import (
	"github.com/sirupsen/logrus"
	"time"
)

type loggerImpl struct {
	log *logrus.Logger
}

func New() Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:          true,
		TimestampFormat:        time.DateTime,
		DisableLevelTruncation: true,
		PadLevelText:           true,
		QuoteEmptyFields:       true,
	})
	return &loggerImpl{
		log: log,
	}
}

func (l *loggerImpl) Info(fields Fields, message string) {
	l.log.WithFields(logrus.Fields(fields)).Infoln(message)
}

func (l *loggerImpl) Error(fields Fields, message string) {
	l.log.WithFields(logrus.Fields(fields)).Errorln(message)
}

func (l *loggerImpl) Fatal(fields Fields, message string) {
	l.log.WithFields(logrus.Fields(fields)).Fatalln(message)
}
