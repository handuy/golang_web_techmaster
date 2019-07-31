package main

import (
	"html/template"
	"net/http"
)

var homeTemplate *template.Template

func homeHandlerFunc(w http.ResponseWriter, r *http.Request) {
	if err := homeTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func main() {
	var err error
	homeTemplate, err = template.ParseFiles("view/index.html", "view/layout/footer.html")
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", homeHandlerFunc)
	http.ListenAndServe(":8080", nil)
}
