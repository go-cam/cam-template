package config

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/base/camBase"
)

// common config
func GetConfig() camBase.AppConfigInterface {
	config := cam.NewConfig()
	config.ComponentDict = map[string]camBase.ComponentConfigInterface{
		"db":      databaseConfig(),
		"console": consoleConfig(),
	}
	return config
}

// database component config
func databaseConfig() camBase.ComponentConfigInterface {
	config := cam.NewDatabaseConfig("mysql", "127.0.0.1", "3306", "go_cam", "root", "root")
	return config
}

// console component config
func consoleConfig() camBase.ComponentConfigInterface {
	config := cam.NewConsoleConfig()
	config.DatabaseDir = "../common/database"
	config.XormTemplateDir = "../common/templates/xorm"
	return config
}
