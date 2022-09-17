package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

type UserController struct {
	BaseController // 继承
}

func (con UserController) Index(c *gin.Context) {
	// c.String(200, "用户列表 ---")
	con.Success(c)
	con.Error(c)
}

func (con UserController) Add(c *gin.Context) {
	//c.String(200, "用户列表-add ---")
	c.HTML(http.StatusOK, "admin/useradd.html", gin.H{})
}

func (con UserController) DoUpload(c *gin.Context) {
	username := c.PostForm("username")
	face, err := c.FormFile("face")
	dst := path.Join("Gin/demo10/static/upload")
	if err == nil {
		c.SaveUploadedFile(face, dst)
	}
	c.JSON(200, gin.H{
		"success":  true,
		"username": username,
		"dst":      dst,
	})

	c.String(200, "执行上传")
}
