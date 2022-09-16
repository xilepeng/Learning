package admin

import "github.com/gin-gonic/gin"

type UserController struct {
	BaseController // 继承
}

func (con UserController) Index(c *gin.Context) {
	// c.String(200, "用户列表 ---")
	con.Success(c)
	con.Error(c)
}

func (con UserController) Add(c *gin.Context) {
	c.String(200, "用户列表-add ---")
}

func (con UserController) Edit(c *gin.Context) {
	c.String(200, "用户列表-edit ---")
}
