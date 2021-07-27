package app

import (
	"github.com/albanybuipe96/bookstore-users-api/controllers/ping"
	"github.com/albanybuipe96/bookstore-users-api/controllers/users"
)

func mapURLS() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}
