package main

import (
	"fmt"
	"github.com/kataras/iris"
)

func homeHandler(ctx iris.Context) {
	ctx.HTML("<h1>Chào mừng đến với Techmaster Việt Nam</h1>")
}

func courseHandler(ctx iris.Context) {
	ctx.HTML("<h1>Đây là trang danh sách khóa học</h1>")
}

func courseDetailHandler(ctx iris.Context) {
	// Đọc dữ liệu từ URL path
	// /khoa-hoc/hoc-Go --> courseName = "hoc-Go"
	// /khoa-hoc/web --> courseName = "web"
	// /khoa-hoc/thuat-toan --> courseName = "thuat-toan"
	courseName := ctx.Params().Get("name")
	htmlString := fmt.Sprintf("<h1>Chi tiết khóa học: %s</h1>", courseName)
	ctx.HTML(htmlString)
}

func courseBeginnerHandler(ctx iris.Context) {
	ctx.HTML("<h1>Khóa học dành cho beginner</h1>")
}

func blogHandler(ctx iris.Context) {
	// Đọc dữ liệu từ query string parameter
	// /bai-viet?category=Mobile&keyword=ios&author=quynh
	//  --> category = "Mobile"
	//  --> keyword = "ios"
	//  --> author = "quynh"
	category := ctx.URLParam("category")
	keyword := ctx.URLParam("keyword")
	author := ctx.URLParam("author")

	htmlCategory := fmt.Sprintf("<h1>Bài viết thuộc danh mục: %s</h1>", category)
	htmlKeyword := fmt.Sprintf("<h1>Từ khóa tìm kiếm: %s</h1>", keyword)
	htmlAuthor := fmt.Sprintf("<h1>Tác giả bài viết: %s</h1>", author)
	
	ctx.HTML(htmlCategory)
	ctx.HTML(htmlKeyword)
	ctx.HTML(htmlAuthor)
}

func main() {
	app := iris.New()
	app.Get("/", homeHandler)
	app.Get("/khoa-hoc", courseHandler)

	// Xử lý các GET request gửi đến: /khoa-hoc/hoc-Go, /khoa-hoc/web, /khoa-hoc/thuat-toan, ...
	app.Get("/khoa-hoc/{name}", courseDetailHandler)

	// Xử lý GET request gửi đến: /khoa-hoc/beginner
	// Không bị mâu thuẫn với API /khoa-hoc/{name} ở trên
	app.Get("/khoa-hoc/beginner", courseBeginnerHandler)

	// Xử lý GET request gửi đến: /bai-viet và đọc dữ liệu trong query string parameter
	app.Get("/bai-viet", blogHandler)

	app.Run(iris.Addr(":8080"))
}