package logging

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Initialized logger
func InitLog() *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(getLogLevel())
	logger.SetFormatter(&logrus.JSONFormatter{})
	return logger
}

// Defines log level from the config
func getLogLevel() logrus.Level {
	switch strings.ToLower(viper.GetString("log.level")) {
	case DebugMode:
		return logrus.DebugLevel
	case InfoMode:
		return logrus.InfoLevel
	case ErrorMode:
		return logrus.ErrorLevel
	case FatalMode:
		return logrus.FatalLevel
	default:
		return logrus.InfoLevel
	}
}
