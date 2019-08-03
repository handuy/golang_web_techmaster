package router

import (
	"golang_web_techmaster/mvc/controller"
	"github.com/kataras/iris"
)

func NewRoutes(c *controller.Controller, api *iris.Application) {
	api.Get("/", c.Hello) 
}