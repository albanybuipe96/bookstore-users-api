package models

import (
	"github.com/albanybuipe96/bookstore-users-api/utils/errors"
	"strings"
)

type User struct {
	Id          int64  `db:"id"`
	FirstName   string `db:"firstname"`
	LastName    string `db:"lastname"`
	Email       string `db:"email"`
	DateCreated string `db:"datecreated"`
}

func (user *User) Validate() *errors.CustomError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
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
	user.DateCreated = usr.DateCreated
}
