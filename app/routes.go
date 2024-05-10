package app

import (
	"github.com/albanybuipe96/bookstore-users-api/handlers"
)

func appRoutes() {
	router.GET("/ping", handlers.Ping)
	v1 := router.Group("/api/v1")
	{
		v1.POST("/users/", handlers.CreateUser)

		v1.GET("/users/:user_id", handlers.GetUser)
		v1.GET("/users/status/:status", handlers.GetUsersByStatus)
		v1.GET("/users/", handlers.GetUsers)

		v1.PATCH("/users/:user_id", handlers.UpdateUser)

		v1.DELETE("/users/:user_id", handlers.DeleteUser)
	}
}
