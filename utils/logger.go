package utils

import (
	"github.com/abdullahjankhan/go_sample/config"
	logInstance "github.com/abdullahjankhan/go_sample/logger"
	"github.com/sirupsen/logrus"
)

type Logger interface {
	Errorf(msg string, args ...interface{})
	Error(err interface{})

	Infof(msg string, args ...interface{})
	Info(msg string)

	Fatalf(msg string, args ...interface{})

	Printf(string, ...interface{})
}

type logger struct {
	loggerInstance *logrus.Logger
}

func NewLogger(conf *config.GlobalConfig) Logger {
	logInstance.Init(conf)
	return &logger{
		loggerInstance: logInstance.Instance(),
	}
}

func (l *logger) Errorf(msg string, args ...interface{}) {
	l.loggerInstance.Errorf(msg, args...)
}

func (l *logger) Error(err interface{}) {
	l.loggerInstance.Error(err)
}

func (l *logger) Infof(msg string, args ...interface{}) {
	l.loggerInstance.Infof(msg, args...)
}

func (l *logger) Fatalf(msg string, args ...interface{}) {
	l.loggerInstance.Fatalf(msg, args...)
}

func (l *logger) Info(msg string) {
	l.loggerInstance.Info(msg)
}

func (l *logger) Printf(key string, value ...interface{}) {
	l.loggerInstance.Infof("%v \t %v", key, value)
}
