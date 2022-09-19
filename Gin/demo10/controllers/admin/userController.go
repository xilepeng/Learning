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

// 单文件上传
func (con UserController) DoUpload(c *gin.Context) {
	username := c.PostForm("username")
	face, err := c.FormFile("face")
	// dst := path.Join("static/upload") // 第一次出错位置，导致没有找到保存路径
	dst := "./static/upload/" + face.Filename
	//dst := path.Join("./static/upload/", face.Filename)
	if err == nil {
		err := c.SaveUploadedFile(face, dst)
		if err != nil {
			return
		}
	}
	c.JSON(200, gin.H{
		"success":  true,
		"username": username,
		"dst":      dst,
	})

	c.String(200, "执行上传")
}

func (con UserController) Edit(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/useredit.html", gin.H{})
}

// 多文件上传
func (con UserController) DoEdit(c *gin.Context) {
	username := c.PostForm("username")

	face1, err1 := c.FormFile("face1")
	//dst := "./static/upload/" + face.Filename
	dst1 := path.Join("./static/upload/", face1.Filename)
	if err1 == nil {
		err := c.SaveUploadedFile(face1, dst1)
		if err != nil {
			return
		}
	}

	face2, err2 := c.FormFile("face2")
	//dst := "./static/upload/" + face.Filename
	dst2 := path.Join("./static/upload/", face2.Filename)
	if err2 == nil {
		err := c.SaveUploadedFile(face2, dst2)
		if err != nil {
			return
		}
	}
	c.JSON(200, gin.H{
		"success":  true,
		"username": username,
		"dst1":     dst1,
		"dst2":     dst2,
	})

	c.String(200, "执行上传")
	//c.String(http.StatusOK, "执行修改")
}
