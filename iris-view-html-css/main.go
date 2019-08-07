package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./view", ".html"))
	
	app.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	})

	app.Get("/khoa-hoc", func(ctx iris.Context) {
		ctx.View("course.html")
	})

	// Trả về file style.css
	app.Get("/css/{file}", func(ctx iris.Context) {
		fileName := ctx.Params().Get("file")
		filePath := "css/" + fileName
		ctx.ServeFile(filePath, true)
	})

	// Trả về file ảnh
	app.Get("/img/{file}", func(ctx iris.Context) {
		fileName := ctx.Params().Get("file")
		filePath := "img/" + fileName
		ctx.ServeFile(filePath, true)
	})
	app.Run(iris.Addr(":8080"))
}