package app

import (
	"github.com/gin-gonic/gin"
	"log"
)

var (
	router = gin.Default()
)

func Start() {
	appRoutes()

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err.Error())
	}
}
