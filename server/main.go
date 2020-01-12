package main

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam-template/server/config"
)

func main() {
	config.LoadConfig()
	cam.App.Run()
}
