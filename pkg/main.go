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
	router.GET("/chat", func(context *gin.Context) {
		context.HTML(http.StatusOK, "chat_homepage.html", gin.H{"title": "homepage"})
	})
	action := router.Group("/action")
	{
		action.GET("/registration", func(context *gin.Context) {
			context.HTML(http.StatusOK, "registration.html", gin.H{"title": "registration"})
		})
		action.GET("/authorization", func(context *gin.Context) {
			context.HTML(http.StatusOK, "authorization.html", gin.H{"title": "authorization"})
		})
	}
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/register", controllers.RegisterUser, func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"username": context.PostForm("username")})
			context.JSON(http.StatusOK, gin.H{"email": context.PostForm("email")})
			context.JSON(http.StatusOK, gin.H{"password": context.PostForm("password")})
		})

		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}
