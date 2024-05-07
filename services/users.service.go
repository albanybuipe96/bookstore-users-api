package services

import (
	"github.com/albanybuipe96/bookstore-users-api/domain/users"
	"time"
)

var db []users.User

var lastId int64 = 0

func GetUserByID(id int64) {

}

func CreateUser(user users.User) (*users.User, error) {
	lastId += 1
	user.Id = lastId
	user.DateCreated = time.Now().UTC().String()
	db = append(db, user)
	return &user, nil
}
