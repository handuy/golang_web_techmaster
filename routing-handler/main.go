package main
import (
	"fmt"
	"net/http"
)

func homeHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Đây là trang chủ</h1>")
}

func courseHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Đây là trang danh sách khóa học</h1>")
}

func main() {
   http.HandleFunc("/", homeHandlerFunc)
   http.HandleFunc("/khoa-hoc", courseHandlerFunc)
   http.ListenAndServe(":8080", nil)
}