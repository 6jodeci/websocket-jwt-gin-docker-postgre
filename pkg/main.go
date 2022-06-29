package main

import (
	"examplechat/controllers"
	"examplechat/database"
	"examplechat/middlewares"
	"net/http"

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
	router.LoadHTMLGlob("../templates/html/*.html")
	router.Static("/css", "../templates/css/")
	router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "chat_homepage.html", gin.H{"title": "homepage"})
	})
	router.GET("/chat/registration", func(context *gin.Context) {
		context.HTML(http.StatusOK, "registration.html", gin.H{"title": "registration"})
	})
	router.GET("/chat/authorization", func(context *gin.Context) {
		context.HTML(http.StatusOK, "authorization.html", gin.H{"title": "authorization"})
	})
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
