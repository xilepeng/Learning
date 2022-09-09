package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	Title   string
	Content string
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	// 前台
	r.GET("/", func(c *gin.Context) {
		//c.String(http.StatusOK, "%v", "你好，gin")
		c.HTML(http.StatusOK, "default/index.html", gin.H{
			"title": "首页",
		})
	})

	r.GET("default/news", func(c *gin.Context) {
		//c.String(http.StatusOK, "%v", "你好，gin")
		news := &Article{
			Title:   "新闻标题",
			Content: "新闻内容",
		}
		c.HTML(http.StatusOK, "default/news.html", gin.H{
			"title": "新闻页面",
			"news":  news,
		})
	})
	// 后台
	r.GET("/admin", func(c *gin.Context) {
		//c.String(http.StatusOK, "%v", "你好，gin")
		c.HTML(http.StatusOK, "admin/index.html", gin.H{
			"title": "后台首页",
		})
	})

	r.GET("/admin/news", func(c *gin.Context) {
		//c.String(http.StatusOK, "%v", "你好，gin")
		news := &Article{
			Title:   "后台新闻标题",
			Content: "后台新闻内容",
		}
		c.HTML(http.StatusOK, "admin/news.html", gin.H{
			"title": "后台新闻页面",
			"news":  news,
		})
	})

	err := r.Run()
	if err != nil {
		return
	}
}
