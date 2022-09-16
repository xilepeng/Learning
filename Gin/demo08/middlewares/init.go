package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func InitMiddleware(c *gin.Context) {
	//判断用户是否登录
	fmt.Println(time.Now())
	fmt.Println(c.Request.URL)
}
