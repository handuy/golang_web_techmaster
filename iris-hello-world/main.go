package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Web server chạy bằng IRIS</h1>")
	})
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}