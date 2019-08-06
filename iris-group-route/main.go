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

func userProfileHandler(ctx iris.Context) {
	ctx.HTML("<h1>Trang profile người dùng</h1>")
}

func userDashboardHandler(ctx iris.Context) {
	ctx.HTML("<h1>Trang dashboard người dùng</h1>")
}

func main() {
	app := iris.New()
	app.Handle("GET", "/", homeHandler)
	app.Handle("GET", "/khoa-hoc", courseHandler)

	app.PartyFunc("/users", func(users iris.Party) {
		// http://localhost:8080/users/profile
		users.Get("/profile", userProfileHandler)
		// http://localhost:8080/users/dashboard
		users.Get("/dashboard", userDashboardHandler)
	})

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
