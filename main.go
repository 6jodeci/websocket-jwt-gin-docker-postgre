package main

import (
	"examplechat/controllers"
	"examplechat/database"
	"examplechat/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	// Init db
	database.Connect("jodeci:6504@tcp(localhost:5432)/private_chat?parseTime=True")
	database.Migrate()
	// Init router
	router := initRouter()
	router.Run(":8000")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}
