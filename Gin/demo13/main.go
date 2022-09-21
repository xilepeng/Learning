package main

import (
	"demo13/models"
	"demo13/routers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"html/template"
)

func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{"UnixToTime": models.UnixToTime})
	r.LoadHTMLGlob("templates/**/*")
	r.Static("/static", "./static")

	// 配置 Session 中间件

	// session 存储到 服务端
	// 创建基于 cookie 的存储引擎，secret 是用于加密的密钥，可以随便写
	//store := cookie.NewStore([]byte("secret"))
	// store 是前面创建的存储引擎，我们可以替换成其他存储引擎
	//r.Use(sessions.Sessions("mysession", store))

	// session 存储到 redis
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	routers.DefaultRoutersInit(r) // Gin 路由文件抽离
	routers.ApiRoutersInit(r)
	routers.AdminRoutersInit(r)

	err := r.Run(":80")
	if err != nil {
		return
	}
}
