package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func indexHander(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "index",
	})
}

func prev(c *gin.Context)  {
	fmt.Println("prev ...")
}

func main() {
	r := gin.Default()
	r.GET("/index", prev,indexHander)
	r.Run()
}
