package services

import (
	"github.com/albanybuipe96/bookstore-users-api/domain/users"
	"github.com/albanybuipe96/bookstore-users-api/utils/errors"
	"time"
)

var db = []users.User{}

var lastId int64 = 0

func CreateUser(user users.User) (*users.User, *errors.CustomError) {
	//return nil, &errors.CustomError{
	//	Message: "User could not be created",
	//	Code:    http.StatusInternalServerError,
	//	Error:   "internal_server_error",
	//}
	lastId += 1
	user.Id = lastId
	user.DateCreated = time.Now().UTC().String()
	db = append(db, user)
	return &user, nil
}

func GetUserByID(id int64) (*[]users.User, *errors.CustomError) {
	return &db, nil
}

/*
{
	"message": "User with given id not found",
	"status": 404,
	"error": "not_found",
}

500: "internal_server_error"
400: "bad_request"

403: "forbidden_resource"
*/
