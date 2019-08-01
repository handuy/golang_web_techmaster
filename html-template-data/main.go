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
	ctx.ViewData("Title", "Đây là trang chủ")
	ctx.ViewData("Name", "Các khóa học mới mở trong tháng")
	ctx.View("index.html")
}