package routers

import "github.com/gin-gonic/gin"

func ApiRoutersInit(r *gin.Engine) {
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
}
