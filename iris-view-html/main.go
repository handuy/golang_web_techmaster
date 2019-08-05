package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./view", ".html"))
	app.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	})
	app.Run(iris.Addr(":8080"))
}