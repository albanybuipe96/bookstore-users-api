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
			"INSERT INTO %s(firstname=?, lastname=?, email=?, status=?, password=?, created=?);",
			query.TableName,
		)
	}
	return fmt.Sprintf(
		"INSERT INTO %s (firstname, lastname, email, status, password, created) VALUES($1, $2, $3, $4, $5, $6) RETURNING *;",
		query.TableName,
	)
}

// Fetch generates an SQL SELECT statement to fetch a user by their ID.
// Returns the SQL statement as a string.
func (query *Query) Fetch() string {
	return fmt.Sprintf("SELECT * FROM %s WHERE id=$1;", query.TableName)
}

// FetchAll generates an SQL SELECT statement to fetch all users from the table.
// Returns the SQL statement as a string.
func (query *Query) FetchAll() string {
	return fmt.Sprintf("SELECT * FROM %s;", query.TableName)
}

// FetchByStatus generates an SQL SELECT statement to fetch all users from the table based on status.
// Returns the SQL statement as a string.
func (query *Query) FetchByStatus() string {
	if query.DbEngine == MySQL {
		return fmt.Sprintf("SELECT * FROM %s WHERE status=?;", query.TableName)
	}
	return fmt.Sprintf("SELECT * FROM %s WHERE status=$1;", query.TableName)
}

// Update generates an SQL UPDATE statement for the given table.
// Returns the SQL statement as a string.
func (query *Query) Update() string {
	if query.DbEngine == MySQL {
		return fmt.Sprintf(
			"UPDATE %s SET firstname=?, "+
				"lastname=?, email=?, status=?, password=?, created=? WHERE id=?;",
			query.TableName,
		)
	}
	return fmt.Sprintf(
		"UPDATE %s SET firstname=$1, lastname=$2, "+
			"email=$3, status=$4, password=$5, created=$6 WHERE id=$7", query.TableName,
	)
}

// Delete generates an SQL DELETE statement for the given table.
// Returns the SQL statement as a string.
func (query *Query) Delete() string {
	if query.DbEngine == MySQL {
		return fmt.Sprintf("DELETE FROM %s WHERE id=?;", query.TableName)
	}
	return fmt.Sprintf("DELETE FROM %s WHERE id=$1", query.TableName)
}
