package main

import (
	"github.com/kataras/iris"
)

func homeHandler(ctx iris.Context) {
	ctx.HTML("<h1>Chào mừng đến với Techmaster Việt Nam</h1>")
}

func main() {
	app := iris.New()
	app.Handle("GET", "/", homeHandler)
	app.HandleDir("/static", "./assets", iris.DirOptions{
		IndexName: "/index.html",
		ShowList: true,
	})
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
