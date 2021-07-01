package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"osaka/startup"
)

func main() {
	err := startup.InitMySQL()
	if err != nil {
		panic(err)
	}

	server := gin.Default()

	server.Use(cors.Default())

	startup.SetupRouters(server)

	server.Run(":8080")
}
