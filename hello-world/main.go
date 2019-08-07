package main
import (
    "fmt"
    "net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
    htmlH1String := "<h1 style='color:red'>Web server chạy bằng Golang</h1>"
    htmlH2String := "<h2 style='color:green; font-size:20px; font-style:Italic;'>Học lập trình back-end</h2>"
    returnHTML := htmlH1String + htmlH2String
    fmt.Fprint(w, returnHTML)
}

func main() {
    http.HandleFunc("/", handlerFunc)
    http.ListenAndServe(":8080", nil)
}