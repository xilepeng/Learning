package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string
	Content string
}

type UserInfo struct {
	UserName string `json:"username" form:"username"`
	PassWord string `json:"password" form:"password"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("demo05/templates/**/*")
	r.Static("/static", "./static")
	// 前台
	// get 请求传值
	r.GET("/", func(c *gin.Context) {
		username := c.Query("username")
		age := c.Query("age")
		page := c.DefaultQuery("page", "1")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"age":      age,
			"page":     page,
		})
		//c.String(http.StatusOK, "%v", "你好，gin")
		//c.HTML(http.StatusOK, "default/index.html", gin.H{
		//	"title": "首页",
		//})
	})
	// post 请求传值，获取表单传过来的数据
	r.GET("/user", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default/user.html", gin.H{})
	})
	r.POST("/doAddUser", func(c *gin.Context) {

		username := c.PostForm("username")
		password := c.PostForm("password")
		age := c.DefaultPostForm("age", "20")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
			"age":      age,
		})
	})

	r.POST("/doAddUser2", func(c *gin.Context) {

		user := &UserInfo{}
		if err := c.ShouldBind(&user); err == nil {
			fmt.Printf("%#v", user)
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusOK, gin.H{"err": err.Error()})
		}
	})

	// 获取 GET POST传递的数据绑定到结构体
	r.GET("/getUser", func(c *gin.Context) {
		user := &UserInfo{}
		if err := c.ShouldBind(&user); err == nil {
			fmt.Printf("%#v", user)
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusOK, gin.H{"err": err.Error()})
		}
	})
	r.GET("/news", func(c *gin.Context) {
		//c.String(http.StatusOK, "%v", "你好，gin")
		news := &Article{
			Title:   "新闻标题",
			Content: "新闻内容",
		}
		c.HTML(http.StatusOK, "default/news.html", gin.H{
			"title": "新闻页面",
			"news":  news,
		})
	})
	// 后台
	r.GET("/admin", func(c *gin.Context) {
		//c.String(http.StatusOK, "%v", "你好，gin")
		c.HTML(http.StatusOK, "admin/index.html", gin.H{
			"title": "后台首页",
		})
	})

	r.GET("/admin/news", func(c *gin.Context) {
		//c.String(http.StatusOK, "%v", "你好，gin")
		news := &Article{
			Title:   "后台新闻标题",
			Content: "后台新闻内容",
		}
		c.HTML(http.StatusOK, "admin/news.html", gin.H{
			"title": "后台新闻页面",
			"news":  news,
		})
	})

	err := r.Run()
	if err != nil {
		return
	}
}
