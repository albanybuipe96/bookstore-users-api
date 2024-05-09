package handlers

import (
	"fmt"
	"github.com/albanybuipe96/bookstore-users-api/domain/models"
	"github.com/albanybuipe96/bookstore-users-api/services"
	"github.com/albanybuipe96/bookstore-users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUser(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		parsingErr := errors.BadRequestError("Invalid json body")
		context.JSON(parsingErr.ReportError())
		return
	}

	result, err := services.CreateUser(&user)
	if err != nil {
		fmt.Println("ERROR HERE", err.Error())
		context.JSON(err.ReportError())
		return
	}
	context.JSON(http.StatusCreated, result)
}

func GetUser(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("user_id"), 10, 64)
	if err != nil {
		er := errors.BadRequestError("Invalid user id")
		context.JSON(er.ReportError())
		return
	}
	result, er := services.GetUserByID(id)
	if er != nil {
		context.JSON(er.ReportError())
		return
	}
	context.JSON(http.StatusOK, result)
}

func GetUsers(context *gin.Context) {
	users, err := services.GetUsers()
	if err != nil {
		context.JSON(err.ReportError())
		return
	}
	context.JSON(http.StatusOK, users)
}
