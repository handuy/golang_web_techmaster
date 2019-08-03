package main

import (
	"github.com/kataras/iris"
	"golang_web_techmaster/mvc/router"
	"golang_web_techmaster/mvc/controller"
) 

func main() {
	app := iris.New()
	c := controller.NewController()
	tmpl := iris.HTML("./view", ".html")
	tmpl.Reload(true)
	app.RegisterView(tmpl)
	router.NewRoutes(c, app)
	app.Run(iris.Addr(":8080")) // defaults to that but you can change it.
}