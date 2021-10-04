package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 1. 定义模板 htllo.tmpl
	// 2. 解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("解析模板错误", err)
	}
	// 3. 渲染模板

	u1 := UserInfo{
		Name:   "X",
		Gender: "男",
		Age:    22,
	}
	m1 := map[string]interface{}{
		"name":   "N",
		"gender": "女",
		"age":    22,
	}
	t.Execute(w, map[string]interface{}{
		"u1": u1,
		"m1": m1,
	})
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("http server start err:", err)
	}
}
