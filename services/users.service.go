package services

import (
	"github.com/albanybuipe96/bookstore-users-api/domain/models"
	"github.com/albanybuipe96/bookstore-users-api/utils/dates"
	"github.com/albanybuipe96/bookstore-users-api/utils/errors"
)

// CreateUser creates a new user in the system.
// It validates the user, sets the creation timestamp, and saves the user to the database.
// If any step fails, it returns an error.
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

func UpdateUser(user *models.User) (*models.User, *errors.CustomError) {
	current, err := GetUserByID(user.Id)
	if err != nil {
		return nil, err
	}
	if err := current.Populate(*user); err != nil {
		return nil, err
	}
	current.CreatedAt = dates.GetFormattedTime()
	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func DeleteUser(user *models.User) (interface{}, *errors.CustomError) {
	current, err := GetUserByID(user.Id)
	if err != nil {
		return "", err
	}
	if err := current.Delete(); err != nil {
		return "", err
	}
	return struct {
		Message string `json:"message"`
	}{
		Message: "User successfully deleted",
	}, nil
}

// GetUserByID retrieves a user by their ID.
// It returns the user if found, or an error if the user does not exist or if there's an issue retrieving the user.
func GetUserByID(id int64) (*models.User, *errors.CustomError) {
	user := &models.User{Id: id}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}

// GetUsers retrieves all users from the system.
// It returns a list of users or an error if there's an issue retrieving the users.
func GetUsers() ([]*models.User, *errors.CustomError) {
	user := &models.User{}
	users, err := user.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetUsers retrieves all users from the system.
// It returns a list of users or an error if there's an issue retrieving the users.
func GetUsersByStatus(status string) ([]*models.User, *errors.CustomError) {
	user := &models.User{}
	users, err := user.GetUserByStatus(status)
	if err != nil {
		return nil, err
	}
	return users, nil
}
