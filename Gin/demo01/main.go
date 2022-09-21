package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.String(200, "%v", "你好，Gin")
	})
	r.GET("/news", func(c *gin.Context) {
		c.String(http.StatusOK, "我是新闻页面")
	})
	r.POST("/add", func(c *gin.Context) {
		c.String(200, "我是post请求，主要用于增加数据")
	})
	r.PUT("/edit", func(c *gin.Context) {
		c.String(200, "我是put请求,用于编辑数据")
	})
	r.DELETE("/delete", func(c *gin.Context) {
		c.String(200, "我是delete请求,删除数据")
	})

	err := r.Run()
	if err != nil {
		return
	} // 监听并在 0.0.0.0:8080 上启动服务
	//r.Run(":8000")
}
