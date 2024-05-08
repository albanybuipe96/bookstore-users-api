package app

import (
	"github.com/albanybuipe96/bookstore-users-api/handlers"
)

func appRoutes() {
	router.GET("/ping", handlers.Ping)
	v1 := router.Group("/api/v1")
	{
		v1.GET("/models/:user_id", handlers.GetUser)
		v1.POST("/models/", handlers.CreateUser)
	}
}
