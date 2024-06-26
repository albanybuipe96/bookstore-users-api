package models

import (
	"log"

	"github.com/albanybuipe96/bookstore-users-api/data/mysql/datasource"
	"github.com/albanybuipe96/bookstore-users-api/data/mysql/queries"
	"github.com/albanybuipe96/bookstore-users-api/utils/errors"
)

// Save saves a user to the database.
// It prepares an insert statement, executes it with the user's details, and returns an error if any step fails.
func (user *User) Save() *errors.CustomError {
	query := queries.Query{TableName: "users", DbEngine: queries.MySQL}
	statement, err := datasource.DbClient.Prepare(query.Insert())

	if err != nil {
		log.Println(err.Error())
		return errors.InternalServerError(err.Error())
	}
	defer statement.Close()

	result, err := statement.Exec(
		user.FirstName,
		user.LastName,
		user.Email,
		user.Status,
		user.Password,
		user.CreatedAt,
	)
	if err != nil {
		return errors.ReportDbError(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return errors.InternalServerError(err.Error())
	}

	user.Id = id

	return nil
}

func (user *User) Update() *errors.CustomError {
	query := queries.Query{TableName: "users", DbEngine: queries.MySQL}
	statement, err := datasource.DbClient.Prepare(query.Update())
	if err != nil {
		log.Println(err.Error())
		return errors.ReportDbError(err)
	}
	defer statement.Close()

	_, err = statement.Exec(
		user.FirstName,
		user.LastName,
		user.Email,
		user.Status,
		user.Password,
		user.CreatedAt,
		user.Id,
	)
	if err != nil {
		log.Println(err.Error())
		return errors.ReportDbError(err)
	}
	return nil
}

// Get retrieves a user from the database by their ID.
// It prepares a fetch statement, executes it with the user's ID, and returns the user or an error if any step fails.
func (user *User) Get() *errors.CustomError {
	query := queries.Query{TableName: "users", DbEngine: queries.MySQL}
	statement, err := datasource.DbClient.Prepare(query.Fetch())
	if err != nil {

		return errors.ReportDbError(err)
	}
	defer statement.Close()

	row := statement.QueryRow(user.Id)

	if err := row.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.CreatedAt,
		&user.Status,
		&user.Password,
	); err != nil {
		log.Println(err.Error())
		return errors.ReportDbError(err)
	}

	return nil
}

// GetAllUsers retrieves all users from the database.
// It prepares a fetch all statement, executes it, and returns a list of users or an error if any step fails.
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
			&user.Status,
			&user.Password,
		)
		if err != nil {
			log.Println(err.Error())
			return nil, errors.InternalServerError(err.Error())
		}
		results = append(results, &user)
	}

	return results, nil
}

func (user *User) GetUserByStatus(status string) ([]*User, *errors.CustomError) {
	query := queries.Query{TableName: "users", DbEngine: queries.MySQL}
	statement, err := datasource.DbClient.Prepare(query.FetchByStatus())
	if err != nil {
		log.Println(err.Error())
		return nil, errors.ReportDbError(err)
	}
	defer statement.Close()

	rows, err := statement.Query(status)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.ReportDbError(err)
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return nil, errors.ReportDbError(err)
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
			&user.Status,
			&user.Password,
		)
		if err != nil {
			log.Println(err.Error())
			return nil, errors.ReportDbError(err)
		}
		results = append(results, &user)
	}

	//if len(results) == 0 {
	//	return nil, errors.NotFoundError("no match found for given status")
	//}

	return results, nil
}

func (user *User) Delete() *errors.CustomError {
	query := queries.Query{TableName: "users", DbEngine: queries.MySQL}
	statement, err := datasource.DbClient.Prepare(query.Delete())
	if err != nil {
		log.Println(err.Error())
		return errors.ReportDbError(err)
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Id); err != nil {
		log.Println(err.Error())
		return errors.ReportDbError(err)
	}
	return nil
}
