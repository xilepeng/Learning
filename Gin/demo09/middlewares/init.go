package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func InitMiddleware(c *gin.Context) {
	//判断用户是否登录
	fmt.Println(time.Now())
	fmt.Println(c.Request.URL)
	// 中间件设置共享数据
	c.Set("username", "席")

	// 在中间件中使用 Goroutine

	// 创建在 goroutine 中使用的副本
	cCp := c.Copy()
	go func() {
		// 用 time.Sleep() 模拟一个长任务。
		time.Sleep(time.Second)

		// 请注意您使用的是复制的上下文 "cCp"，这一点很重要
		log.Println("Done! in path " + cCp.Request.URL.Path)
	}()
}
