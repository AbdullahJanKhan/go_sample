package service

import (
	"github.com/abdullahjankhan/go_sample/config"
	"github.com/abdullahjankhan/go_sample/logger"
	"github.com/sirupsen/logrus"
)

type LoggerService interface {
	GetInstance() *logrus.Logger
}

type loggerService struct {
	logger *logrus.Logger
}

func NewLoggerService(conf *config.GlobalConfig) LoggerService {
	logger.Init(conf)
	return &loggerService{
		logger: logger.Instance(),
	}

}

func (log *loggerService) GetInstance() *logrus.Logger {
	return log.logger
}
