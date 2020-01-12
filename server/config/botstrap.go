package config

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam-template/common/config"
	"github.com/go-cam/cam-template/server/controllers"
)

func LoadConfig() {
	cam.App.AddConfig(config.GetCommonConfig())
	cam.App.AddConfig(GetServerConfig())

	loadRouteConfig()
}

func loadRouteConfig() {
	router := cam.App.GetRouter()
	router.Register(new(controllers.HelloController))
}
