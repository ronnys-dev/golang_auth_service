package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ronnys-dev/golang_auth_service/database"
	"github.com/ronnys-dev/golang_auth_service/handlers"
	"github.com/ronnys-dev/golang_auth_service/middlewares"
)

func main() {
	config := database.Config{}
	config.FillConfig()
	database.Connect(config.DatabaseUrl())
	database.Migrate()

	router := initRouters()
	router.Run(":8000")
}

func initRouters() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", handlers.GenerateToken)
		api.POST("/user/register", handlers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", handlers.Ping)
		}
	}
	return router
}