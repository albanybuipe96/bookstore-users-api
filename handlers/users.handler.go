package handlers

import (
	"encoding/json"
	"github.com/albanybuipe96/bookstore-users-api/domain/users"
	"github.com/albanybuipe96/bookstore-users-api/services"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func CreateUser(context *gin.Context) {
	var user users.User

	bytes, err := io.ReadAll(context.Request.Body)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = json.Unmarshal(bytes, &user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.CreateUser(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, result)
}

func GetUser(context *gin.Context) {

}

func FindUser(context *gin.Context) {}
