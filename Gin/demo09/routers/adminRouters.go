package routers

import (
	"Gin/demo09/controllers/admin"
	"Gin/demo09/middlewares"
	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string
	Content string
}

func AdminRoutersInit(r *gin.Engine) {
	// 后台
	adminRouter := r.Group("/admin", middlewares.InitMiddleware) //全局添加中间件
	{
		adminRouter.GET("/", admin.IndexController{}.Index)

		adminRouter.GET("/user", admin.UserController{}.Index) // 自定义控制器
		adminRouter.GET("/user/add", admin.UserController{}.Add)
		adminRouter.GET("/user/edit", admin.UserController{}.Edit)

		adminRouter.GET("/article", admin.ArticleController{}.Index)
		adminRouter.GET("/article/add", admin.ArticleController{}.Add)
		adminRouter.GET("/article/edit", admin.ArticleController{}.Edit)

		// adminRouter.GET("/news", func(c *gin.Context) {
		// 	//c.String(http.StatusOK, "%v", "你好，gin")
		// 	news := &Article{
		// 		Title:   "后台新闻标题",
		// 		Content: "后台新闻内容",
		// 	}
		// 	c.HTML(http.StatusOK, "admin/news.html", gin.H{
		// 		"title": "后台新闻页面",
		// 		"news":  news,
		// 	})
		// })
	}
}
