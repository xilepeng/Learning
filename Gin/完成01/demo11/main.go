package main

import "github.com/gin-gonic/gin"

func main()  {
	r := gin.Default()
	r.GET("/web", func(c *gin.Context) {
		//name := c.Query("query")
		//name := c.DefaultQuery("query", "somebody")
		name, ok := c.GetQuery("query")
		if !ok {
			name = "somebody"
		}
		c.JSON(200,gin.H{
			"name":name,
		})
	})
	r.Run(":8080")
}
