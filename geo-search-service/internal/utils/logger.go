package utils

import (
	"github.com/sirupsen/logrus"
)

const APP_NAME = "geo-search"

var appName string = APP_NAME

func BuildLogger(app string) {
	appName = app
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func newLogger() *logrus.Entry {
	if appName == "" {
		appName = APP_NAME
	}
	return logrus.WithField("app_name", appName)
}

func LogDebug(args ...interface{}) {
	newLogger().Debug(args...)
}

func LogInfo(args ...interface{}) {
	newLogger().Info(args...)
}

func LogInfof(format string, args ...interface{}) {
	newLogger().Infof(format, args...)
}

func LogWarn(args ...interface{}) {
	newLogger().Warn(args...)
}

func LogError(args ...interface{}) {
	newLogger().Error(args...)
}

func LogErrorf(format string, args ...interface{}) {
	newLogger().Errorf(format, args...)
}

func LogPanic(args ...interface{}) {
	newLogger().Panic(args...)
}

func LogPanicf(format string, args ...interface{}) {
	newLogger().Panicf(format, args...)
}

func LogFatal(args ...interface{}) {
	newLogger().Fatal(args...)
}

func LogFatalf(format string, args ...interface{}) {
	newLogger().Fatalf(format, args...)
}
