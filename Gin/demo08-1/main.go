package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string
	Content string
}

// 中间件：匹配路由之前~匹配路由之后，所执行的一系列操作
// 路由中间件,在路由前执行，可实现权限判断
func initMiddleware(c *gin.Context) {
	start := time.Now().UnixNano()
	fmt.Println("1-我是中间件,在路由前执行，可实现权限判断")

	c.Next() // 调用该请求的剩余处理程序

	fmt.Println("2-我是中间件,在路由前执行，可实现权限判断")
	end := time.Now().UnixNano()
	fmt.Println("程序执行时间：", end-start)
}

//1-我是中间件,在路由前执行，可实现权限判断
//这是一个首页
//2-我是中间件,在路由前执行，可实现权限判断

func main() {
	r := gin.Default()
	// r.LoadHTMLGlob("demo08-1/templates/**/*")
	r.Static("/static", "./static")
	// 前台
	r.GET("/", initMiddleware, func(c *gin.Context) {
		fmt.Println("这是一个首页")
		c.String(http.StatusOK, "%v", "gin首页")
		time.Sleep(time.Second)
		// c.HTML(http.StatusOK, "default/index.html", gin.H{
		// 	"title": "首页",
		// })
	})

	r.GET("/news", initMiddleware, func(c *gin.Context) {
		c.String(http.StatusOK, "%v", "gin新闻")
		// news := &Article{
		// 	Title:   "新闻标题",
		// 	Content: "新闻内容",
		// }
		// c.HTML(http.StatusOK, "default/news.html", gin.H{
		// 	"title": "新闻页面",
		// 	"news":  news,
		// })
	})

	// 后台
	// r.GET("/admin", func(c *gin.Context) {
	// 	//c.String(http.StatusOK, "%v", "你好，gin")
	// 	c.HTML(http.StatusOK, "admin/index.html", gin.H{
	// 		"title": "后台首页",
	// 	})
	// })

	// r.GET("/admin/user", func(c *gin.Context) {
	// 	//c.String(http.StatusOK, "%v", "你好，gin")
	// 	c.HTML(http.StatusOK, "admin/index.html", gin.H{
	// 		"title": "用户列表",
	// 	})
	// })

	// r.GET("/admin/news", func(c *gin.Context) {
	// 	//c.String(http.StatusOK, "%v", "你好，gin")
	// 	news := &Article{
	// 		Title:   "后台新闻标题",
	// 		Content: "后台新闻内容",
	// 	}
	// 	c.HTML(http.StatusOK, "admin/news.html", gin.H{
	// 		"title": "后台新闻页面",
	// 		"news":  news,
	// 	})
	// })

	err := r.Run()
	if err != nil {
		return
	}
}
