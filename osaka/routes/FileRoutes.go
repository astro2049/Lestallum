package routes

import (
	"github.com/gin-gonic/gin"
	"osaka/service"
)

func InitFileRouter(Router *gin.RouterGroup) {
	FileRoutes := Router.Group("/file")
	{
		FileRoutes.GET("/sts", service.HandOutCOSCredential)
		FileRoutes.POST("/register", service.RegisterFile)
		FileRoutes.GET("/info", service.GetFileInfo)
	}
}
