package models

import (
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

// Validate checks if the user's email is valid.
// It trims whitespace and converts the email to lowercase.
// If the email is empty or does not contain an '@' symbol, it returns an error.
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

// Populate copies the values from one User instance to another.
// It's used to transfer data between instances of the User struct.
func (user *User) Populate(usr User) {
	user.Id = usr.Id
	user.FirstName = usr.FirstName
	user.LastName = usr.LastName
	user.Email = usr.Email
	user.CreatedAt = usr.CreatedAt
}
