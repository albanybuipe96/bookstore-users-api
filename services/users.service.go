package services

import (
	"fmt"
	"github.com/albanybuipe96/bookstore-users-api/domain/users"
	"github.com/google/uuid"
	"os"
)

func GetUserByID(id string) users.User {
	fmt.Fprintf(os.Stdout, "Finding user with given id: %v\n", id)
	return users.User{
		Id:       uuid.New().String(),
		Username: "albanybuipe",
		Email:    "albanybuipe@nexus.com",
		Password: "Password@1",
	}
}
