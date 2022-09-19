package routers

import (
	"demo10-2/controllers/api"

	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine) {
	apiRouter := r.Group("/api")
	{
		apiRouter.GET("/", api.ApiController{}.Index)
		apiRouter.GET("/userlist", api.ApiController{}.UserList)
		apiRouter.GET("/plist", api.ApiController{}.PList)
	}
}
