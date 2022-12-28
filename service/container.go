package service

import (
	"github.com/abdullahjankhan/go_sample/config"
	"github.com/abdullahjankhan/go_sample/lib"
	"github.com/abdullahjankhan/go_sample/repository/db_name"
	"github.com/abdullahjankhan/go_sample/utils"
)

type Container struct {
	Logger              utils.Logger
	GlobalConfigService GlobleConfigService
	HashingService      HashingService
	JWTService          JWTService
	LoggerService       LoggerService
	SampleService       SampleService
}

func NewService() *Container {
	globalConfig := config.GetConfig()
	globalConfigService := NewGbeConfigService()
	_ = lib.GetHttpClient()
	logger := utils.NewLogger(globalConfig)
	jwtService := NewJWTService(globalConfigService)
	hashingService := NewHashingService(nil, 0, 0, 0, 0, 0)
	loggerService := NewLoggerService(globalConfig)

	db := db_name.SharedStore()
	sampleService := NewSampleService(db)

	return &Container{
		Logger:              logger,
		GlobalConfigService: globalConfigService,
		HashingService:      hashingService,
		JWTService:          jwtService,
		LoggerService:       loggerService,
		SampleService:       sampleService,
	}

}
