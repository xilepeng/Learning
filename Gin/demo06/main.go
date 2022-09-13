package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string
	Content string
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("demo06/templates/**/*")
	r.Static("/static", "./static")

	defaultRouters := r.Group("/")
	{
		// 前台
		defaultRouters.GET("/", func(c *gin.Context) {
			//c.String(http.StatusOK, "%v", "你好，gin")
			//c.HTML(http.StatusOK, "/index.html", gin.H{
			//	"title": "首页",
			//})
			c.String(200, "%v", "首页")
		})

		defaultRouters.GET("/news", func(c *gin.Context) {
			c.String(http.StatusOK, "%v", "新闻")
			//news := &Article{
			//	Title:   "新闻标题",
			//	Content: "新闻内容",
			//}
			//c.HTML(http.StatusOK, "/news.html", gin.H{
			//	"title": "新闻页面",
			//	"news":  news,
			//})
		})
	}

	apiRouter := r.Group("/api")
	{
		apiRouter.GET("/", func(c *gin.Context) {
			c.String(200, "我是一个api接口")
		})
		apiRouter.GET("/userlist", func(c *gin.Context) {
			c.String(200, "我是一个api接口 userlist")
		})
		apiRouter.GET("/plist", func(c *gin.Context) {
			c.String(200, "我是一个api接口 plist")
		})
	}

	// 后台
	adminRouter := r.Group("/admin")
	{
		adminRouter.GET("/", func(c *gin.Context) {
			//c.String(http.StatusOK, "%v", "你好，gin")
			c.HTML(http.StatusOK, "admin/index.html", gin.H{
				"title": "后台首页",
			})
		})

		adminRouter.GET("/news", func(c *gin.Context) {
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

		adminRouter.GET("/article", func(c *gin.Context) {
			c.String(http.StatusOK, "%v", "新闻列表")
			//c.HTML(http.StatusOK, "admin/index.html", gin.H{
			//	"title": "新闻列表",
			//})
		})
	}

	err := r.Run()
	if err != nil {
		return
	}
}
