package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default() // used in url_mappings.go
)

func StartApplication() {
	mapURLS()
	router.Run(":8080")
}
