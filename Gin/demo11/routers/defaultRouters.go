package routers

import (
	"demo11/controllers/xn"

	"github.com/gin-gonic/gin"
)

// Gin 路由文件抽离
func DefaultRoutersInit(r *gin.Engine) {

	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", xn.DefaultController{}.Index)
		defaultRouters.GET("/news", xn.DefaultController{}.News)
		// 前台
		// defaultRouters.GET("/", func(c *gin.Context) {
		// 	//c.String(http.StatusOK, "%v", "你好，gin")
		// 	//c.HTML(http.StatusOK, "/index.html", gin.H{
		// 	//	"title": "首页",
		// 	//})
		// 	c.String(200, "%v", "首页")
		// })

		// defaultRouters.GET("/news", func(c *gin.Context) {
		// 	c.String(http.StatusOK, "%v", "新闻")
		// 	//news := &Article{
		// 	//	Title:   "新闻标题",
		// 	//	Content: "新闻内容",
		// 	//}
		// 	//c.HTML(http.StatusOK, "/news.html", gin.H{
		// 	//	"title": "新闻页面",
		// 	//	"news":  news,
		// 	//})
		// })

	}
}
