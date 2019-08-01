package main

import (
	"github.com/kataras/iris"
) 

func main() {
	app := iris.New()
	// Đăng kí các file HTML trong thư mục view
	tmpl := iris.HTML("./view", ".html")
	// Mỗi lần chỉnh sửa file HTML thì chỉ cần refresh lại trang
	tmpl.Reload(true)
	app.RegisterView(tmpl)
	app.Get("/", hi)
	app.Run(iris.Addr(":8080")) // defaults to that but you can change it.
}

func hi(ctx iris.Context) {
	// Truyền dữ liệu vào 2 biến Title và Name và pass 2 biến này vào file index.html
	ctx.ViewData("Name", "Java")
	ctx.ViewData("Age", 10)
	ctx.ViewData("IsNew", true)
	ctx.View("index.html")
}