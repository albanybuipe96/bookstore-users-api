package models

import (
	"net/mail"
	"strings"

	"github.com/albanybuipe96/bookstore-users-api/utils/errors"
)

// User represents a user in the system.
// It includes fields for ID, first name, last name, email, and creation timestamp.
type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	CreatedAt string `json:"created"`
}

func NewUser() *User {
	return &User{}
}

// Validate checks if the user's email is valid.
// It trims whitespace and converts the email to lowercase.
// If the email is empty or does not contain an '@' symbol, it returns an error.
func (user *User) Validate() *errors.CustomError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return errors.BadRequestError("invalid email address")
	}
	return nil
}

// Populate copies the values from one User instance to another.
// It's used to transfer data between instances of the User struct.
func (user *User) Populate(usr User) *errors.CustomError {
	user.Id = usr.Id
	if usr.FirstName != "" {
		user.FirstName = usr.FirstName
	}
	if usr.LastName != "" {
		user.LastName = usr.LastName
	}
	if usr.Email != "" {
		if err := usr.Validate(); err != nil {
			return errors.BadRequestError(err.Error())
		}
		user.Email = usr.Email
	}
	if usr.CreatedAt != "" {
		user.CreatedAt = usr.CreatedAt
	}
	return nil
}
