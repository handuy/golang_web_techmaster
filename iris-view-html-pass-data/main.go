package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()
	tmpl := iris.HTML("./view", ".html")
	app.RegisterView(tmpl)
	// Mỗi lần chỉnh sửa file HTML thì chỉ cần refresh lại trang
	tmpl.Reload(true)
	
	app.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	})
	app.Get("/detail/{name}", func(ctx iris.Context) {
		// Đọc dữ liệu từ URL path
		heroName := ctx.Params().Get("name")

		// Đọc dữ liệu từ query string parameter
		place := ctx.URLParam("place")
		email := ctx.URLParam("email")
		ctx.ViewData("name", heroName)
		ctx.ViewData("place", place)
		ctx.ViewData("email", email)
		ctx.View("detail.html")
	})

	app.Run(iris.Addr(":8080"))
}