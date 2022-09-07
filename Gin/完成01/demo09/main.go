package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

func main() {
	r := gin.Default()
	r.Static("/xxx","./statics")
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	//解析模板
	r.LoadHTMLGlob("templates/**/*")
	//渲染模板
	r.GET("posts/index", func(c *gin.Context) {
		c.HTML(200, "posts/index.html", gin.H{
			"title": "posts/index",
		})
	})
	r.GET("users/index", func(c *gin.Context) {
		c.HTML(200, "users/index.html", gin.H{
			"title": "<a href='https://liwenzhou.com'>李文周的博客</a>",
		})
	})
	//启动服务
	r.Run(":8080")
}
