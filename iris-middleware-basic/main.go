package main

import (
	"log"
	"time"

	"github.com/kataras/iris"
)

func before(ctx iris.Context) {
	log.Printf("Thời gian nhận request: %s", time.Now())
	ctx.Next()
}

func after(ctx iris.Context) {
	result := ctx.Values().Get("main")
	log.Println(result)
	log.Println("Xử lý xong request")
}

func handler(ctx iris.Context) {
	ctx.HTML("<h1>Tạo middleware với IRIS framework</h1>")
	ctx.Values().Set("main", "received from main handler")
	ctx.Next()
}

func main() {
	app := iris.New()
	app.Get("/", before, handler, after)
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
