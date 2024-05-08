package models

import (
	"fmt"
	"github.com/albanybuipe96/bookstore-users-api/data/mysql/datasource"
	"github.com/albanybuipe96/bookstore-users-api/data/mysql/queries"
	"github.com/albanybuipe96/bookstore-users-api/utils/errors"
	"log"
)

var users = make(map[int64]*User)

func (user *User) Save() *errors.CustomError {
	//query := queries.Query{TableName: "users", DbEngine: queries.MySQL}
	//statement, err := datasource.DbClient.Prepare(query.Insert())

	fmt.Println("Got here 1")
	statement, err := datasource.DbClient.Prepare(
		"INSERT INTO users(firstname, lastname, email, created) VALUES(?, ?, ?, ?);",
	)
	fmt.Println("Got here")
	if err != nil {
		fmt.Println("Error here")
		log.Println(err.Error())
		return errors.InternalServerError(err.Error())
	}
	defer statement.Close()

	fmt.Println("Got here 2")

	result, err := statement.Exec(
		user.FirstName, user.LastName, user.Email, user.CreatedAt,
	)
	if err != nil {
		log.Println(err.Error())
		return errors.InternalServerError(err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return errors.InternalServerError(err.Error())
	}

	user.Id = id

	return nil
}

func (user *User) Get() *errors.CustomError {
	query := queries.Query{TableName: "users", DbEngine: queries.PostgreSQL}
	statement, err := datasource.DbClient.Prepare(query.Fetch())
	if err != nil {
		log.Println(err.Error())
		return errors.InternalServerError(err.Error())
	}
	defer statement.Close()

	_, err = statement.Query(user.Id)
	if err != nil {
		log.Println(err.Error())
		return errors.InternalServerError(err.Error())
	}

	return nil
}
