package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("templates/*.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	err = tmpl.ExecuteTemplate(w, "index.tmpl", nil)
	if err != nil {
		fmt.Println("render template failed, err:", err)
		return
	}
}

func main() {
	http.HandleFunc("/index", index)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("http server start err:", err)
	}
}
