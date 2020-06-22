package controllers

import "github.com/go-cam/cam"

type HelloController struct {
	cam.Controller
}

func (ctrl *HelloController) GetDefaultActionName() string {
	return "Cam"
}

func (ctrl *HelloController) Cam() {
	ctrl.GetContext().Write([]byte("Welcome to cam framework!"))
}
