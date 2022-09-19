package xn

import (
	"demo10-2/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

type DefaultController struct{}

func (con DefaultController) Index(c *gin.Context) {
	// c.String(200, "我是首页")
	fmt.Println(models.UnixToTime(1662774111))
	c.HTML(200, "default/index.html", gin.H{
		"date": 1662774111,
	})
}

func (con DefaultController) News(c *gin.Context) {
	c.String(200, "我是新闻")
}
