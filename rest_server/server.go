package restserver

import (
	"io/ioutil"
	"time"

	"github.com/abdullahjankhan/go_sample/controller"
	"github.com/abdullahjankhan/go_sample/service"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func StartServer() {
	service := service.NewService()
	controller := controller.NewController(service)
	middleware := NewMiddleware(service.JWTService)
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	router := gin.Default()
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           int(12 * time.Hour),
	}))
	router.ForwardedByClientIP = true
	NewRoutes(router, middleware, controller).InitRoutes()

	err := router.Run(service.GlobalConfigService.GetConfig().RestServer.Addr)
	if err != nil {
		panic(err)
	}
	service.Logger.Info("#==================================#")
	service.Logger.Infof("#===== Server Running At: %v ====#", service.GlobalConfigService.GetConfig().RestServer.Addr)
	service.Logger.Info("#==================================#")
}
