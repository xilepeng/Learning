package routers

import (
	"demo12/controllers/admin"
	"demo12/middlewares"

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
		// 1. 单文件上传
		adminRouter.GET("/user/add", admin.UserController{}.Add)
		adminRouter.POST("/user/doUpload", admin.UserController{}.DoUpload)
		// 2. 多文件上传
		adminRouter.GET("/user/edit", admin.UserController{}.Edit)
		adminRouter.POST("/user/doEdit", admin.UserController{}.DoEdit)
		// 3. 相同名文件上传
		adminRouter.GET("/user/addsame", admin.UserController{}.AddSame)
		adminRouter.POST("/user/doUploadSame", admin.UserController{}.DoUploadSame)

		// 文章
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
