package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func main()  {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		//方法一：使用map
		c.JSON(http.StatusOK,gin.H{ "message":"hello"})
	})
	//方法二：结构体
	type msg struct {
		Name string `json:name`
	}
	r.GET("morejson", func(c *gin.Context) {

		data := msg{
			Name: "X",
		}
		c.JSON(http.StatusOK, data)
	})
	r.Run(":8080")

}
