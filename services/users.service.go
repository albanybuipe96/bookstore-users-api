package services

import (
	"github.com/albanybuipe96/bookstore-users-api/domain/models"
	"github.com/albanybuipe96/bookstore-users-api/utils/errors"
	"time"
)

var db = []models.User{}

var lastId int64 = 0

func CreateUser(user models.User) (*models.User, *errors.CustomError) {
	return nil, errors.InternalServerError("User could not be created")
	lastId += 1
	user.Id = lastId
	user.DateCreated = time.Now().UTC().String()
	db = append(db, user)
	return &user, nil
}

func GetUserByID(id int64) (*[]models.User, *errors.CustomError) {
	return &db, nil
}
