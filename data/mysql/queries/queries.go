package queries

import "fmt"

// Engine represents the type of database engine being used.
// It is an enumeration with values for PostgreSQL and MySQL.
type Engine int

const (
	PostgreSQL Engine = iota
	MySQL
)

// Query represents a database query.
// It contains the table name and the database engine type.
type Query struct {
	TableName string
	DbEngine  Engine
}

// Insert generates an SQL INSERT statement for the given table.
// It uses the database engine type to format the statement correctly.
// Returns the SQL statement as a string.
func (query *Query) Insert() string {
	if query.DbEngine == MySQL {
		return fmt.Sprintf(
			"INSERT INTO %s(firstname, lastname, email, created) VALUES(?, ?, ?, ?);",
			query.TableName,
		)
	}
	return fmt.Sprintf(
		"INSERT INTO %s (firstname, lastname, email, created) VALUES($1, $2, $3, $4);",
		query.TableName,
	)
}

// Fetch generates an SQL SELECT statement to fetch a user by their ID.
// It uses the table name to construct the statement.
// Returns the SQL statement as a string.
func (query *Query) Fetch() string {
	return fmt.Sprintf("SELECT * FROM %s WHERE id=?;", query.TableName)
}

// FetchAll generates an SQL SELECT statement to fetch all users from the table.
// It uses the table name to construct the statement.
// Returns the SQL statement as a string.
func (query *Query) FetchAll() string {
	return fmt.Sprintf("SELECT * FROM %s;", query.TableName)
}
