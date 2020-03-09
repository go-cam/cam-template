package config

import (
	"app/server/controllers"
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/base/camBase"
)

func GetConfig() camBase.AppConfigInterface {
	config := cam.NewConfig()
	config.ComponentDict = map[string]camBase.ComponentConfigInterface{
		"http": httpConfig(),
		"ws":   websocketConfig(),
	}
	return config
}

// http component config
func httpConfig() camBase.ComponentConfigInterface {
	config := cam.NewHttpConfig(8800)
	config.Register(&controllers.HelloController{})
	return config
}

// websocket component config
func websocketConfig() camBase.ComponentConfigInterface {
	config := cam.NewWebsocketConfig(9800)
	config.Register(&controllers.HelloController{})
	return config
}
