package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping is a handler function for the ping endpoint.
// It responds with a JSON object containing a message "pong".
// This endpoint is typically used to check if the server is running.
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
