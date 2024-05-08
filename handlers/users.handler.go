package handlers

import (
	"github.com/albanybuipe96/bookstore-users-api/domain/models"
	"github.com/albanybuipe96/bookstore-users-api/services"
	"github.com/albanybuipe96/bookstore-users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		parsingErr := errors.BadRequestError("Invalid json body")
		context.JSON(parsingErr.ReportError())
		return
	}

	result, err := services.CreateUser(user)
	if err != nil {
		context.JSON(err.ReportError())
		return
	}
	context.JSON(http.StatusCreated, result)
}

func GetUser(context *gin.Context) {
	result, err := services.GetUserByID(context.GetInt64("user_id"))
	if err != nil {
		context.JSON(err.ReportError())
		return
	}
	context.JSON(http.StatusOK, result)
}
