package main

import (
	"github.com/kataras/iris"
	"fmt"
)

func homeHandler(ctx iris.Context) {
	ctx.HTML("<h1>Chào mừng đến với Techmaster Việt Nam</h1>")
}

func serveFileHandler(ctx iris.Context) {
	folder := ctx.Params().Get("folder")
	file := ctx.Params().Get("file")
	// đường dẫn trỏ đến file trong thư mục assets
	filePath := fmt.Sprintf("./assets/%s/%s", folder, file)
	ctx.ServeFile(filePath, false)
}

func main() {
	app := iris.New()
	app.Get("/", homeHandler)
	app.Get("/static/{folder}/{file}", serveFileHandler)
	
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}