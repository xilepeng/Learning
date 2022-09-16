package admin

import "github.com/gin-gonic/gin"

type BaseController struct{}

func (con BaseController) Success(c *gin.Context) {
	c.String(200, "成功")
}

func (con BaseController) Error(c *gin.Context) {
	c.String(200, "失败")
}
