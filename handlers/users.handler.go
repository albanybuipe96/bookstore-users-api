package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/albanybuipe96/bookstore-users-api/domain/models"
	"github.com/albanybuipe96/bookstore-users-api/services"
	"github.com/albanybuipe96/bookstore-users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(context *gin.Context) {
	user := models.NewUser()

	if err := context.ShouldBindJSON(&user); err != nil {
		parsingErr := errors.BadRequestError("Invalid json body")
		context.JSON(parsingErr.ReportError())
		return
	}

	result, err := services.CreateUser(user)
	if err != nil {
		fmt.Println(err.Error())
		context.JSON(err.ReportError())
		return
	}
	response := models.Response{
		Data:    result,
		Message: "User added successfully",
		Code:    http.StatusCreated,
	}
	context.JSON(http.StatusCreated, response)
}

func GetUser(context *gin.Context) {
	id, err := getId(context)
	if err != nil {
		context.JSON(err.ReportError())
		return
	}
	result, er := services.GetUserByID(id)
	if er != nil {
		context.JSON(er.ReportError())
		return
	}
	response := models.Response{
		Data:    result,
		Message: "success",
		Code:    http.StatusOK,
	}
	context.JSON(http.StatusOK, response)
}

func UpdateUser(context *gin.Context) {
	id, err := getId(context)
	if err != nil {
		context.JSON(err.ReportError())
		return
	}
	user := models.NewUser()

	if err := context.ShouldBindJSON(&user); err != nil {
		fmt.Println(err.Error())
		parsingErr := errors.BadRequestError("Invalid json body")
		context.JSON(parsingErr.ReportError())
		return
	}

	user.Id = id
	result, err := services.UpdateUser(user)
	if err != nil {
		fmt.Println(err.Error())
		context.JSON(err.ReportError())
		return
	}
	response := models.Response{
		Data:    result,
		Message: "success",
		Code:    http.StatusOK,
	}
	context.JSON(http.StatusOK, response)

}

func DeleteUser(context *gin.Context) {
	id, err := getId(context)
	if err != nil {
		context.JSON(err.ReportError())
		return
	}
	user := models.NewUser()
	user.Id = id
	result, err := services.DeleteUser(user)
	if err != nil {
		fmt.Println(err.Error())
		context.JSON(err.ReportError())
		return
	}
	response := models.Response{
		Data:    result,
		Message: "success",
		Code:    http.StatusOK,
	}
	context.JSON(http.StatusOK, response)
}

func GetUsers(context *gin.Context) {
	users, err := services.GetUsers()
	if err != nil {
		context.JSON(err.ReportError())
		return
	}
	response := models.Response{
		Data:    users,
		Message: "success",
		Code:    http.StatusOK,
	}
	context.JSON(http.StatusOK, response)
}

func getId(context *gin.Context) (int64, *errors.CustomError) {
	id, err := strconv.ParseInt(context.Param("user_id"), 10, 64)
	if err != nil {
		return 0, errors.BadRequestError("Invalid user id")
	}
	return id, nil
}
