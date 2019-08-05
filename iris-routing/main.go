package main

import (
	"fmt"
	"github.com/kataras/iris"
)

type BlogInfo struct {
	Id string 
	Name string
}

func homeHandler(ctx iris.Context) {
	ctx.HTML("<h1>Chào mừng đến với Techmaster Việt Nam</h1>")
}

func courseHandler(ctx iris.Context) {
	ctx.HTML("<h1>Đây là trang danh sách khóa học</h1>")
}

func courseDetailHandler(ctx iris.Context) {
	courseName := ctx.Params().Get("name")
	result := fmt.Sprintf("<h1>Đây là khóa học %s</h1>", courseName)
	ctx.HTML(result)
}

func blogHandler(ctx iris.Context) {
	ctx.HTML("<h1>Đây là trang danh sách bài viết</h1>")
}

func blogDetailHandler(ctx iris.Context) {
	blogId := ctx.Params().Get("id")
	blogTitle := ctx.Params().Get("name")
	var result = BlogInfo{
		Id: blogId,
		Name: blogTitle,
	}
	ctx.JSON(result)
}

func notFoundHandler(ctx iris.Context) {
	ctx.HTML("<h1>Oops! Trang bạn tìm không tồn tại</h1>")
}

func main() {
	app := iris.New()
	app.Get("/", homeHandler)

	app.Get("/khoa-hoc", courseHandler)
	app.Get("/khoa-hoc/{name}", courseDetailHandler)

	app.Get("/bai-viet", blogHandler)
	app.Get("/bai-viet/{id}/{name}", blogDetailHandler)

	app.OnErrorCode(iris.StatusNotFound, notFoundHandler)
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}