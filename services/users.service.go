package services

import (
	"github.com/albanybuipe96/bookstore-users-api/domain/models"
	"github.com/albanybuipe96/bookstore-users-api/utils/dates"
	"github.com/albanybuipe96/bookstore-users-api/utils/errors"
)

func CreateUser(user *models.User) (*models.User, *errors.CustomError) {
	if err := user.Validate(); err != nil {
		return nil, errors.BadRequestError(err.Error())
	}

	user.CreatedAt = dates.GetFormattedTime()

	if err := user.Save(); err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByID(id int64) (*models.User, *errors.CustomError) {
	user := &models.User{Id: id}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}

func GetUsers() ([]*models.User, *errors.CustomError) {
	user := &models.User{}
	users, err := user.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}
