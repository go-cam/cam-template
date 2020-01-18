package config

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/core/models"
)

func GetCommonConfig() *cam.Config {
	config := cam.NewConfig()
	config.AppConfig = appConfig()
	return config
}

func appConfig() *models.AppConfig {
	appConfig := cam.NewAppConfig()

	return appConfig
}
