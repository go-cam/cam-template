package config

import (
	"app/common/config"
	"app/server/controllers"
	_ "app/server/database/migrations"
	"github.com/go-cam/cam"
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
