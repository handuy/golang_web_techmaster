package main

import (
	"github.com/kataras/iris"
)

func homeHandler(ctx iris.Context) {
	ctx.HTML("<h1>Chào mừng đến với Techmaster Việt Nam</h1>")
}

func courseHandler(ctx iris.Context) {
	ctx.HTML("<h1>Danh sách khóa học của Techmaster Việt Nam</h1>")
}

func main() {
	app := iris.New()
	app.Handle("GET", "/", homeHandler)
	app.Handle("GET", "/khoa-hoc", courseHandler)
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
