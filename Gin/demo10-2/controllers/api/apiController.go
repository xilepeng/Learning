package api

import (
	"github.com/gin-gonic/gin"
)

type ApiController struct{}

func (con ApiController) Index(c *gin.Context) {
	c.String(200, "我是一个api接口")
}

func (con ApiController) UserList(c *gin.Context) {
	c.String(200, "我是一个api接口-userlist")
}

func (con ApiController) PList(c *gin.Context) {
	c.String(200, "我是一个api接口-plist")
}
