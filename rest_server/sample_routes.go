package restserver

import (
	"github.com/abdullahjankhan/go_sample/controller"
	"github.com/abdullahjankhan/go_sample/models"
	"github.com/gin-gonic/gin"
)

type SampleRoutes interface {
	InitRoutes()
}

type sampleRoutes struct {
	router     *gin.Engine
	middleware Middleware
	controller *controller.Controller
}

func NewSampleRoutes(router *gin.Engine, middleware Middleware, controller *controller.Controller) SampleRoutes {
	return &sampleRoutes{
		router:     router,
		middleware: middleware,
		controller: controller,
	}
}

func (sr *sampleRoutes) InitRoutes() {
	routes := sr.router.Group("/v1", AttachBody[models.SampleRequest]())
	{
		routes.POST("/api/sample/create", sr.controller.SampleController.SampleFunc)
	}
}
