package controller

import (
	"github.com/kataras/iris"
	"golang_web_techmaster/mvc/model"
)

func (c *Controller) Hello(ctx iris.Context) {
	ctx.ViewData("courses", model.Courses)
	ctx.View("index.html")
}