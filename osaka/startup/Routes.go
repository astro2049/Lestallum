package startup

import (
	"github.com/gin-gonic/gin"
	"osaka/routes"
)

func SetupRouters(server *gin.Engine) {
	Group := server.Group("/api")
	{
		routes.InitFileRouter(Group)
	}
}
