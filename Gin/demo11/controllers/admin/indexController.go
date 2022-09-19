package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type IndexController struct{}

func (con IndexController) Index(c *gin.Context) {
	username, _ := c.Get("username")
	fmt.Println(username)
	// 类型断言
	if v, ok := username.(string); ok { // 空接口转换为 string
		c.String(200, "用户列表---"+v)
	} else {
		c.String(200, "用户列表---获取用户失败")
	}
}
