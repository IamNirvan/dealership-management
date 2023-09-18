package util

import (
	"github.com/sirupsen/logrus"
)

func LogDebugMsg(message string) {
	logrus.Debug(message)
}

func LogInfoMsg(message string) {
	logrus.Info(message)
}

func LogTraceMsg(message string) {
	logrus.Trace(message)
}

func LogErrorMsg(message string) {
	logrus.Error(message)
}

func LogFatalMsg(message string) {
	logrus.Fatal(message)
}
