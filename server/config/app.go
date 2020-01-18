package config

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/core/base"
)

func GetServerConfig() *cam.Config {
	config := cam.NewConfig()
	config.ComponentDict = map[string]base.ConfigComponentInterface{
		"http": httpConfig(),
		"db":   dbConfig(),
	}
	return config
}

func httpConfig() base.ConfigComponentInterface {
	return cam.NewConfigHttpServer(10080).SetSessionName("cam-template")
}

func dbConfig() base.ConfigComponentInterface {
	driverName := cam.App.GetEvn("DB_DRIVER_NAME")
	host := cam.App.GetEvn("DB_HOST")
	port := cam.App.GetEvn("DB_PORT")
	name := cam.App.GetEvn("DB_NAME")
	username := cam.App.GetEvn("DB_USERNAME")
	password := cam.App.GetEvn("DB_PASSWORD")

	return cam.NewConfigDatabase(driverName, host, port, name, username, password)
}
