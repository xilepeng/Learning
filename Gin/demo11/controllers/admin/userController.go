package admin

import (
	"demo11/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
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

// // 1. 单文件上传
//
//	func (con UserController) DoUpload(c *gin.Context) {
//		username := c.PostForm("username")
//		face, err := c.FormFile("face")
//		// dst := path.Join("static/upload") // 第一次出错位置，导致没有找到保存路径
//		dst := "./static/upload/" + face.Filename
//		//dst := path.Join("./static/upload/", face.Filename)
//		if err == nil {
//			err := c.SaveUploadedFile(face, dst)
//			if err != nil {
//				return
//			}
//		}
//		c.JSON(200, gin.H{
//			"success":  true,
//			"username": username,
//			"dst":      dst,
//		})
//
//		c.String(200, "执行上传")
//	}

// 1. 单文件上传 （按日期存）
func (con UserController) DoUpload(c *gin.Context) {
	username := c.PostForm("username")
	// a.获取上传文件
	file, err := c.FormFile("face")
	// dst := path.Join("static/upload") // 第一次出错位置，导致没有找到保存路径
	//dst := "./static/upload/" + file.Filename
	//dst := path.Join("./static/upload/", face.Filename)
	if err == nil {
		// b.获取后缀名，判断类型是否正确 .jpg .png .gif .jpeg
		extName := path.Ext(file.Filename)
		allowExtMap := map[string]bool{".jpg": true, ".png": true, ".gif": true, ".jpeg": true}
		if _, ok := allowExtMap[extName]; !ok {
			c.String(200, "上传的文件类型不合法")
			return
		}
		// c.创建文件保存的目录 static/upload/20060102
		day := models.GetDay()
		//dir := path.Join("./static/upload/", day)
		dir := "./static/upload/" + day
		if err := os.MkdirAll(dir, 0777); err != nil {
			fmt.Println(err)
			c.String(200, "MkdirAll 失败")
			return
		}
		// d.生成文件名和文件保存目录

		fileName := strconv.FormatInt(models.GetUnix(), 10) + extName

		// e.执行上传
		dst := path.Join(dir, fileName)
		err := c.SaveUploadedFile(file, dst)
		if err != nil {
			return
		}
	}
	c.JSON(200, gin.H{
		"success":  true,
		"username": username,
		// "dst":      dst,
	})

	c.String(200, "执行上传")
}

// 2. 多文件上传
func (con UserController) Edit(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/useredit.html", gin.H{})
}

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

// 3.上传相同名文件

func (con UserController) AddSame(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/useraddsame.html", gin.H{})
}

// 多文件上传
func (con UserController) DoUploadSame(c *gin.Context) {
	username := c.PostForm("username")

	// Multipart form
	form, _ := c.MultipartForm()
	files := form.File["face[]"]

	for _, file := range files {
		log.Println(file.Filename)
		dst := path.Join("./static/upload/", file.Filename)
		// 上传文件至指定目录
		c.SaveUploadedFile(file, dst)
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))

	//face1, err1 := c.FormFile("face1")
	////dst := "./static/upload/" + face.Filename
	//dst1 := path.Join("./static/upload/", face1.Filename)
	//if err1 == nil {
	//	err := c.SaveUploadedFile(face1, dst1)
	//	if err != nil {
	//		return
	//	}
	//}
	//
	//face2, err2 := c.FormFile("face2")
	////dst := "./static/upload/" + face.Filename
	//dst2 := path.Join("./static/upload/", face2.Filename)
	//if err2 == nil {
	//	err := c.SaveUploadedFile(face2, dst2)
	//	if err != nil {
	//		return
	//	}
	//}
	c.JSON(200, gin.H{
		"success":  true,
		"username": username,
		//"dst":     dst,
		//"dst2":     dst2,
	})

	c.String(200, "执行上传")
	//c.String(http.StatusOK, "执行修改")
}
