package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path"
)

func main() {
	router := gin.Default()
	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		//dst := "./" + file.Filename
		dst := path.Join("./static/upload/", file.Filename)
		err := c.SaveUploadedFile(file, dst)
		if err != nil {
			return
		}
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded", file.Filename))
	})
	err := router.Run()
	if err != nil {
		return
	}
}
