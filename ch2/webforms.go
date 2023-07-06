package main

import (
	"log"
	"net/http"
	"text/template"
)

func Home(writer http.ResponseWriter, reader *http.Request) {
	var template_html *template.Template
	template_html = template.Must(template.ParseFiles("ch2/main.html"))
	_ = template_html.Execute(writer, nil)
}

func main() {
	log.Println("Server started on: http://localhost:8000")
	http.HandleFunc("/", Home)
	_ = http.ListenAndServe(":8000", nil)
}
