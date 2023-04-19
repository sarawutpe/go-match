package router

import (
	"main/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {

	apiv1 := r.Group("/api")
	{
		apiv1.POST("/test", service.CreateUser)
		apiv1.POST("/upload", service.Upload)
		apiv1.POST("/remove/:id", service.RemoveFile)

		apiv1.GET("/jwt", service.JWT)

	}

}
