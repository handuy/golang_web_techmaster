package main

import (
	"time"

	"github.com/kataras/iris"
)

func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Thời gian nhận request: %s", time.Now())
	ctx.Next()
}

func main() {
	app := iris.New()
	app.Use(myMiddleware)
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Tạo middleware với IRIS framework</h1>")
	})
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
