package main

import (
	"app/server/config"
	"github.com/go-cam/cam"
)

func main() {
	config.LoadConfig()
	cam.App.Run()
}
