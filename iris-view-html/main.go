package main

import (
	"fmt"

	"github.com/kataras/iris"
)

func homeHandler(ctx iris.Context) {
	// Render file index.html trong thư mục ./view để trả về cho client
	ctx.View("index.html")
}

func blogHandler(ctx iris.Context) {
	// Render file index.html trong thư mục ./view/blog để trả về cho client
	ctx.View("blog/index.html")
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
	// Đăng kí thư mục chứa file HTML
	// ./view là đường dẫn đến thư mục view
	// .html là định dạng các file HTML trả về
	tmpl := iris.HTML("./views", ".html")
	app.RegisterView(tmpl)
	tmpl.Reload(true)

	app.Get("/", homeHandler)

	app.Get("/bai-viet", blogHandler)

	app.Get("/static/{folder}/{file}", serveFileHandler)

	app.Run(iris.Addr(":8080"))
}
