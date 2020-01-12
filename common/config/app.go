package config

import (
	"github.com/go-cam/cam"
)

func GetCommonConfig() *cam.Config {
	config := cam.NewConfig()
	config.AppConfig = cam.NewAppConfig()
	return config
}
