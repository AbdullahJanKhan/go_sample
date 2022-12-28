package restserver

import (
	"github.com/abdullahjankhan/go_sample/controller"
	"github.com/gin-gonic/gin"
)

type RestRoutes struct {
	SampleRoutes SampleRoutes
}

func NewRoutes(router *gin.Engine, middleware Middleware, controller *controller.Controller) RestRoutes {
	sampleRoutes := NewSampleRoutes(router, middleware, controller)

	return RestRoutes{
		SampleRoutes: sampleRoutes,
	}
}

func (r RestRoutes) InitRoutes() {
	r.SampleRoutes.InitRoutes()
}
