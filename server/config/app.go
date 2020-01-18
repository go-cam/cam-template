package config

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/core/base"
	_ "github.com/go-sql-driver/mysql"
)

func GetServerConfig() *cam.Config {
	config := cam.NewConfig()
	config.ComponentDict = map[string]base.ConfigComponentInterface{
		"http":    httpServerConfig(),
		"db":      databaseConfig(),
		"console": consoleConfig(),
	}
	return config
}

// HttpServer config
func httpServerConfig() base.ConfigComponentInterface {
	return cam.NewHttpServerConfig(10080).SetSessionName("cam-template")
}

// Database config
func databaseConfig() base.ConfigComponentInterface {
	driverName := cam.App.GetEvn("DB_DRIVER_NAME")
	host := cam.App.GetEvn("DB_HOST")
	port := cam.App.GetEvn("DB_PORT")
	name := cam.App.GetEvn("DB_NAME")
	username := cam.App.GetEvn("DB_USERNAME")
	password := cam.App.GetEvn("DB_PASSWORD")

	return cam.NewDatabaseConfig(driverName, host, port, name, username, password)
}

// Console config
func consoleConfig() base.ConfigComponentInterface {
	return cam.NewConsoleConfig()
}
