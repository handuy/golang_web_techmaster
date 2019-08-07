package main

import "github.com/kataras/iris"

type BlogDetail struct {
	Id string
	Title string
	Content string
	Author string
}

var blogList = []BlogDetail{
	{
		Id: "1",
		Title: "Học Golang cấp tốc: Struct&Method",
		Content: "Học Golang cấp tốc: Struct&Method",
		Author: "Trịnh Thúy",
	},
	{
		Id: "2",
		Title: "Học Golang cấp tốc: Interface",
		Content: "Học Golang cấp tốc: Interface",
		Author: "Nguyễn Long",
	},
	{
		Id: "3",
		Title: "SQL căn bản: SELECT",
		Content: "SQL căn bản: SELECT",
		Author: "Nguyễn Duy",
	},
	{
		Id: "4",
		Title: "SQL căn bản: Điều kiện WHERE",
		Content: "SQL căn bản: Điều kiện WHERE",
		Author: "Nguyễn Duy",
	},
	{
		Id: "5",
		Title: "Lập trình hướng đối tượng trong Java",
		Content: "Lập trình hướng đối tượng trong Java",
		Author: "Robin Huy",
	},
}

func main() {
	app := iris.New()
	tmpl := iris.HTML("./view", ".html")
	app.RegisterView(tmpl)
	// Mỗi lần chỉnh sửa file HTML thì chỉ cần refresh lại trang
	tmpl.Reload(true)
	
	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("blogList", blogList)
		ctx.View("index.html")
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