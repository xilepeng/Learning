package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 2. 解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("ParseFiles failed:", err)
	}
	// 3. 渲染模板
	name := "mojo"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Println("ParseFiles failed:", err)
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("http server start failed:", err)
		return
	}
}
