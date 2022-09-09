package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("demo02/templates/*")
	r.GET("/", func(c *gin.Context) {
		c.String(200, "%v", "首页")
	})
	r.GET("/json1", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"success": true,
			"msg":     "你好,gin",
		})
	})
	r.GET("/json2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"msg":     "你好,gin, json2",
		})
	})

	r.GET("/json3", func(c *gin.Context) {
		a := &Article{Title: "野蛮进化", Desc: "必须把自己逼的更紧", Content: "意志力"}
		c.JSON(200, a)
	})
	// 响应 jsonp请求，解决跨域问题
	r.GET("/jsonp", func(c *gin.Context) {
		a := &Article{Title: "野蛮进化-jsonp", Desc: "必须把自己逼的更紧", Content: "意志力"}
		c.JSONP(200, a)
	})
	// http://localhost:8080/jsonp?callback=x
	// x (
	// 	{
	// 	"title": "野蛮进化-jsonp",
	// 	"desc": "必须把自己逼的更紧",
	// 	"content": "意志力"
	// 	}
	// 	)
	r.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"success": true,
			"msg":     "你好gin, 我是xml数据",
		})
	})

	r.GET("/news", func(c *gin.Context) {
		c.HTML(http.StatusOK, "news.html", gin.H{
			"title": "我是后台数据",
		})
	})
	r.GET("/goods", func(c *gin.Context) {
		c.HTML(http.StatusOK, "goods.html", gin.H{
			"title": "我是商品数据",
			"price": 20,
		})
	})

	err := r.Run()
	if err != nil {
		return
	}
}
