package main

import "github.com/gin-gonic/gin"
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html","./index.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})

	r.POST("/login", func(c *gin.Context) {
		//username := c.PostForm("username")
		//password := c.PostForm("password")
		username := c.DefaultPostForm("username", "somebody")
		password := c.DefaultPostForm("password", "******")
		c.HTML(200, "index.html", gin.H{
			"Name":username,
			"Password":password,
		})
	})
	r.Run(":8080")
}
