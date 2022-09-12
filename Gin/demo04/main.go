package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 时间戳转换成日期
func UnixToTime(timestamp int) string {
	fmt.Println("时间戳：", timestamp)
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func main() {
	r := gin.Default()
	// 自定义模板功能,放在路由下，模版前
	r.SetFuncMap(template.FuncMap{"UnixToTime": UnixToTime})
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "首页标题",
			"date":  1662774111,
		})
	})
	err := r.Run()
	if err != nil {
		return
	}
}
