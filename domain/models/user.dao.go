package models

import (
	"fmt"
	"github.com/albanybuipe96/bookstore-users-api/utils/dates"
	"github.com/albanybuipe96/bookstore-users-api/utils/errors"
)

var users = make(map[int64]*User)

func (user *User) Save() *errors.CustomError {
	current := users[user.Id]

	for _, usr := range users {
		if usr.Email == user.Email {
			return errors.BadRequestError("Email already taken")
		}
	}

	if current != nil {
		return errors.BadRequestError("User already exists")
	}

	user.DateCreated = dates.GetFormattedTime()
	users[user.Id] = user
	return nil
}

func (user *User) Get() *errors.CustomError {
	result := users[user.Id]
	if result == nil {
		return errors.NotFoundError(
			fmt.Sprintf("User with id %v not found", user.Id),
		)
	}
	user.Populate(*result)
	return nil
}
