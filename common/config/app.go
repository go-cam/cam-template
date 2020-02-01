package config

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/core/camModels"
)

func GetCommonConfig() *cam.Config {
	config := cam.NewConfig()
	config.AppConfig = appConfig()
	return config
}

func appConfig() *camModels.AppConfig {
	appConfig := cam.NewAppConfig()

	return appConfig
}
