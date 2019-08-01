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
	type CourseInfo struct {
		Name string
		Price int
	}

	var courses = []CourseInfo{
		{
			Name: "Xây dựng web báng hàng bằng Golang",
			Price: 2400000,
		},
		{
			Name: "Web cơ bản HTML5, CSS3, Javascript",
			Price: 600000,
		},
		{
			Name: "Thiết kế UI/UX",
			Price: 2400000,
		},
		{
			Name: "Python phân tích dữ liệu",
			Price: 6000000,
		},
		{
			Name: "Lộ trình đào tạo mobile 6 tháng cam kết việc làm",
			Price: 2400000,
		},
	}

	ctx.ViewData("courses", courses)
	ctx.View("index.html")
}