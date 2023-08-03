package config

import (
	"github.com/sirupsen/logrus"
)

type Log struct {
	Level string `mapstructure:"API_LOG_LEVEL"`
}

func (l *Log) New() *logrus.Entry {
	logLevel, err := logrus.ParseLevel(l.Level)
	if err != nil {
		logLevel = logrus.InfoLevel
	}
	logger := logrus.New()
	logger.SetLevel(logLevel)

	log := logrus.NewEntry(logger)

	return log
}
