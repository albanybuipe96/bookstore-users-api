package models

import (
	"fmt"
	"github.com/albanybuipe96/bookstore-users-api/data/mysql/datasource"
	"github.com/albanybuipe96/bookstore-users-api/data/mysql/queries"
	"github.com/albanybuipe96/bookstore-users-api/utils/errors"
	"log"
)

var users = make(map[int64]*User)

// TODO: 11:19:32 CONTINUE FROM HERE

func (user *User) Save() *errors.CustomError {
	query := queries.Query{TableName: "users", DbEngine: queries.MySQL}
	statement, err := datasource.DbClient.Prepare(query.Insert())

	if err != nil {
		log.Println(err.Error())
		return errors.InternalServerError(err.Error())
	}
	defer statement.Close()

	result, err2 := statement.Exec(
		user.FirstName, user.LastName, user.Email, user.CreatedAt,
	)
	if err2 != nil {
		fmt.Println("ERROR HERE")
		log.Println(err2.Error())
		return errors.InternalServerError(err2.Error())
	}

	id, err3 := result.LastInsertId()
	if err3 != nil {
		log.Println(err3.Error())
		return errors.InternalServerError(err3.Error())
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
