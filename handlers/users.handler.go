package handlers

import (
	"github.com/albanybuipe96/bookstore-users-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(context *gin.Context) {
	context.String(http.StatusNotImplemented, "Not Implemented")
}

func GetUser(context *gin.Context) {
	id := context.Param("user_id")
	user := services.GetUserByID(id)
	context.JSON(http.StatusOK, user)
}

func FindUser(context *gin.Context) {}
