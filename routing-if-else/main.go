package main
import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    if r.URL.Path == "/" {
        fmt.Fprint(w, "<h1>Đây là trang chủ</h1>")
    } else if r.URL.Path == "/khoa-hoc" {
        fmt.Fprint(w, "<h1>Đây là trang danh sách khóa học</h1>")
    } else {
        fmt.Fprint(w, "<h1>Không tìm thấy trang</h1>")
    }
}

func main() {
   http.HandleFunc("/", handlerFunc)
   http.ListenAndServe(":8080", nil)
}