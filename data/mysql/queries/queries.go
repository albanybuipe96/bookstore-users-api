package queries

import "fmt"

const (
	PostgreSQL Engine = iota
	MySQL
)

type Query struct {
	TableName string
	DbEngine  Engine
}

type Engine int

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

func (query *Query) Fetch() string {
	return fmt.Sprintf("SELECT * FROM %s WHERE id=?;", query.TableName)
}

func (query *Query) FetchAll() string {
	return fmt.Sprintf("SELECT * FROM %s;", query.TableName)
}
