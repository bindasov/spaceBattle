package logger

import (
	log "github.com/sirupsen/logrus"
)

type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}

func InitLogger(config *Config) (Logger, error) {
	logger := log.New()

	level, err := log.ParseLevel(config.Level)
	if err != nil {
		// use info Level by default
		level = log.InfoLevel
	}
	logger.SetLevel(level)

	return logger, nil
}
