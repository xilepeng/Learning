package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// RESTful API
	r.GET("book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "get",
		})
	})

	r.POST("book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "post",
		})
	})

	r.PUT("book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "put",
		})
	})

	r.DELETE("book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "delete",
		})
	})

	r.Run()
}
