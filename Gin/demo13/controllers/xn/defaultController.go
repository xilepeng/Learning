package xn

import (
	"demo13/models"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type DefaultController struct{}

func (con DefaultController) Index(c *gin.Context) {
	// 设置 session
	session := sessions.Default(c)
	session.Set("username", "乔丹")
	session.Save() // 设置 session 必须调用

	c.String(200, "我是首页")
	fmt.Println(models.UnixToTime(1662774111))

	// 设置 cookie, 实现不同页面数据共享,要放在渲染HTML之前
	//c.SetCookie("username", "我来防乔丹", 3600, "/", ".hfbpw.com", false, true)

	// 过期时间演示
	//c.SetCookie("hobby", "吃饭、上班、学习、睡觉、", 5, "/", "localhost", false, true)

	//c.HTML(200, "default/index.html", gin.H{
	//	"date": 1662774111,
	//})

}

func (con DefaultController) News(c *gin.Context) {

	// 获取 cookie
	// username, _ := c.Cookie("username")
	// hobby, _ := c.Cookie("hobby")
	// c.String(200, "新闻页获取 cookie:"+username+"\n")
	// c.String(200, "新闻页获取 cookie:"+hobby)

	// 获取 session
	session := sessions.Default(c)
	username := session.Get("username")
	c.String(200, "获取 session：username=%v\n", username)
}

func (con DefaultController) Shop(c *gin.Context) {

	// 获取 cookie
	username, _ := c.Cookie("username")

	hobby, _ := c.Cookie("hobby")
	c.String(200, "商品页获取 cookie:"+username+"\n")

	c.String(200, "商品页获取 cookie:"+hobby)
}

// 删除 Cookie
func (con DefaultController) DeleteCookie(c *gin.Context) {
	// 设置 cookie, 实现不同页面数据共享,要放在渲染之前
	c.SetCookie("username", "我来防乔丹", -1, "/", "localhost", false, true)
	c.String(200, "删除成功！")
}
