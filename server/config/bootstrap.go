package config

import (
	"app/common/config"
	"github.com/go-cam/cam"
)

func Bootstrap() {
	// load common's config
	cam.App.AddConfig(config.GetConfig())
	// load module's config
	cam.App.AddConfig(GetConfig())
}
