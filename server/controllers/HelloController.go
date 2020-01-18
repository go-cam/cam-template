package controllers

import (
	"github.com/go-cam/cam"
)

type HelloController struct {
	cam.BaseController
}

func (c *HelloController) HelloWorld() {
	c.Write([]byte("hello world!"))
}
