package main

import (
	"Gin/demo08/routers" // 08 写成07导致包导入错误，无法访问：404 page not found

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//r.LoadHTMLGlob("demo08/templates/**/*")
	//r.Static("/static", "./static")

	routers.DefaultRoutersInit(r) // Gin 路由文件抽离
	routers.ApiRoutersInit(r)
	routers.AdminRoutersInit(r)

	err := r.Run()
	if err != nil {
		return
	}
}
