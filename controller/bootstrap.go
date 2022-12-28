package controller

import (
	"github.com/abdullahjankhan/go_sample/config"
	"github.com/abdullahjankhan/go_sample/models"
	"github.com/abdullahjankhan/go_sample/service"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	Conf             config.GlobalConfig
	SampleController SampleController
}

func NewController(service *service.Container) *Controller {
	sampleController := NewSampleController(service.SampleService, service.Logger)
	return &Controller{
		Conf:             *service.GlobalConfigService.GetConfig(),
		SampleController: sampleController,
	}
}

func GetBody[BodyType any](c *gin.Context) BodyType {
	return c.MustGet(models.RequestKey).(BodyType)
}
