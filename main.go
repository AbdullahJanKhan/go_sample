package main

import (
	"github.com/abdullahjankhan/go_sample/config"
	grpc "github.com/abdullahjankhan/go_sample/grpc_server"
	"github.com/abdullahjankhan/go_sample/kafka/reader/sample_reader/engine"
	"github.com/abdullahjankhan/go_sample/kafka/workers"
	server "github.com/abdullahjankhan/go_sample/rest_server"
	"github.com/abdullahjankhan/go_sample/service"
)

func main() {

	config.SetConfFilePath("./")
	service := service.NewService()

	service.Logger.Info("#==================================#")
	service.Logger.Info("#===========Starting Server =======#")
	service.Logger.Info("#==================================#")

	go server.StartServer(service)

	service.Logger.Info("#==================================#")
	service.Logger.Infof("#===========Rest Server @localhost:%v =======#", service.GlobalConfigService.GetConfig().RestServer.Addr)
	service.Logger.Info("#==================================#")

	/*
	* Initiate Kafka Engine
	 */

	service.Logger.Info("#==================================#")
	service.Logger.Infof("#=========== Starting Kafka Reader =======#")
	service.Logger.Info("#==================================#")

	workers.
		NewSampleLogReader(
			engine.
				NewKafkaLogReader(
					"newSampleLogReader",
					service.GlobalConfigService.GetConfig().Kafka.Brokers,
					service.Logger,
				),
			service.Logger).Start()
	service.Logger.Info("#===============Kafka Reader Started===================#")

	/*
	* Initiate GRPC Server
	 */

	service.Logger.Info("#==================================#")
	service.Logger.Infof("#=========== Starting gRPC Server =======#")
	service.Logger.Info("#==================================#")

	grpc.StartServer(service)

	service.Logger.Infof("#=========== gRPC Server Started =======#")

	select {}
}
