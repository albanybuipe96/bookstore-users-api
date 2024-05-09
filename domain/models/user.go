package models

import (
	"github.com/albanybuipe96/bookstore-users-api/utils/errors"
	"strings"
)

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	CreatedAt string `json:"created"`
}

func (user *User) Validate() *errors.CustomError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" || !strings.Contains(user.Email, "@") {
		return &errors.CustomError{
			Message: "invalid email address",
			Code:    500,
			Reason:  "internal_server_error",
		}
	}
	return nil
}

func (user *User) Populate(usr User) {
	user.Id = usr.Id
	user.FirstName = usr.FirstName
	user.LastName = usr.LastName
	user.Email = usr.Email
	user.CreatedAt = usr.CreatedAt
}
