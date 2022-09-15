package xn

import (
	"github.com/gin-gonic/gin"
)

type DefaultController struct{}

func (con DefaultController) Index(c *gin.Context) {
	c.String(200, "我是首页")
}

func (con DefaultController) News(c *gin.Context) {
	c.String(200, "我是新闻")
}
