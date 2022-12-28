package main

import (
	"github.com/abdullahjankhan/go_sample/config"
	"github.com/abdullahjankhan/go_sample/logger"
	server "github.com/abdullahjankhan/go_sample/rest_server"
)

func main() {

	config.SetConfFilePath("./")
	globalConfig := config.GetConfig()
	logger.Init(globalConfig)

	logger.Instance().Info("#==================================#")
	logger.Instance().Info("#===========Starting Server =======#")
	logger.Instance().Info("#==================================#")

	go server.StartServer()

	select {}
}
