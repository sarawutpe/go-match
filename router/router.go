package router

import (
	"main/middleware"
	"main/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {

	authRoutes := r.Group("/api", middleware.AuthJWTMiddleware())
	{
		authRoutes.GET("/user", service.GetUser)
		authRoutes.POST("/test", service.CreateUser)
		authRoutes.POST("/upload", service.Upload)
		authRoutes.POST("/remove/:id", service.RemoveFile)

	}

	publicRoutes := r.Group("/api")
	{
		publicRoutes.GET("/jwt", service.JWT)
		publicRoutes.GET("/refresh-token/:refresh-token", service.RefreshToken)
	}

}
