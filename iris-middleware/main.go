package main

import (
	"log"
	"time"

	"github.com/kataras/iris"
)

func before(ctx iris.Context) {
	log.Printf("Tên đường dẫn: %s. Thời gian nhận request: %s", ctx.Path(), time.Now())
	ctx.Next()
}

func mainHandler(ctx iris.Context) {
	ctx.HTML("<h1>Trang chủ</h1>")
}

func courseHandler(ctx iris.Context) {
	ctx.HTML("<h1>Trang danh sách khóa học</h1>")
}

func main() {
	app := iris.New()
	// Đăng kí hàm before làm middleware cho cả 2 route / và /khoa-hoc
	app.Use(before)
	app.Get("/", mainHandler)
	app.Get("/khoa-hoc", courseHandler)
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
