package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	Title   string
	Content string
}

func AdminRoutersInit(r *gin.Engine) {
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
}
