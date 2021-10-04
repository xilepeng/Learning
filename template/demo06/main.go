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

func f1(w http.ResponseWriter, r *http.Request) {

	k := func(name string) (string, error) {
		return name + "年轻帅气", nil
	}
	// 1.定义模板

	// 创建一个名为 f 的模板对象
	t := template.New("f.tmpl")
	//告诉模板引擎，添加了自定义函数 kua
	t.Funcs(template.FuncMap{
		"kua": k,
	})
	// 2.解析模板
	_, err := t.ParseFiles("./f.tmpl")
	if err != nil {
		fmt.Println("解析模板错误 err:", err)
	}
	// 3.渲染模板
	name := "X"
	t.Execute(w, name)
}

func tmplDemo(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	user := UserInfo{
		Name:   "小王子",
		Gender: "男",
		Age:    18,
	}
	tmpl.Execute(w, user)
}

func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/tmpl", tmplDemo)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("http server start err:", err)
	}
}
