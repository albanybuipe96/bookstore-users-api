package models

import (
	"github.com/albanybuipe96/bookstore-users-api/data/mysql/datasource"
	"github.com/albanybuipe96/bookstore-users-api/data/mysql/queries"
	"github.com/albanybuipe96/bookstore-users-api/utils/errors"
	"log"
	"strings"
)

func (user *User) Save() *errors.CustomError {
	query := queries.Query{TableName: "users", DbEngine: queries.MySQL}
	statement, err := datasource.DbClient.Prepare(query.Insert())

	if err != nil {
		log.Println(err.Error())
		return errors.InternalServerError(err.Error())
	}
	defer statement.Close()

	result, err := statement.Exec(
		user.FirstName, user.LastName, user.Email, user.CreatedAt,
	)
	if err != nil {
		log.Println(err.Error())

		if strings.Contains(err.Error(), "email_UNIQUE") {
			return errors.BadRequestError("email already taken")
		}

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
	query := queries.Query{TableName: "users", DbEngine: queries.MySQL}
	statement, err := datasource.DbClient.Prepare(query.Fetch())
	if err != nil {
		log.Println(err.Error())
		return errors.InternalServerError(err.Error())
	}
	defer statement.Close()

	row := statement.QueryRow(user.Id)

	if err := row.Scan(
		&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt,
	); err != nil {
		log.Println(err.Error())
		if strings.Contains(err.Error(), "no rows in result set") {
			return errors.NotFoundError("user not found")
		}
		return errors.NotFoundError(err.Error())
	}

	return nil
}

func (user *User) GetAllUsers() ([]*User, *errors.CustomError) {
	query := queries.Query{TableName: "users", DbEngine: queries.MySQL}
	statement, err := datasource.DbClient.Prepare(query.FetchAll())
	if err != nil {
		log.Println(err.Error())
		return nil, errors.InternalServerError(err.Error())
	}
	defer statement.Close()

	rows, err := statement.Query()
	if err != nil {
		log.Println(err.Error())
		return nil, errors.InternalServerError(err.Error())
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Println(err.Error())
		return nil, errors.InternalServerError(err.Error())
	}
	results := make([]*User, 0, len(columns))
	for rows.Next() {
		var user User
		err = rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.CreatedAt,
		)
		if err != nil {
			log.Println(err.Error())
			return nil, errors.InternalServerError(err.Error())
		}
		results = append(results, &user)
	}

	return results, nil
}
