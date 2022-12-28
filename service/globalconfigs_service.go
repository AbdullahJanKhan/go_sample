package service

import (
	"github.com/abdullahjankhan/go_sample/config"
)

type GlobleConfigService interface {
	GetConfig() *config.GlobalConfig
}

type globleConfigService struct {
}

func NewGbeConfigService() GlobleConfigService {
	return &globleConfigService{}
}

func (c *globleConfigService) GetConfig() *config.GlobalConfig {

	return config.GetConfig()

}
