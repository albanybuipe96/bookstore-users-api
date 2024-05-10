package datasource

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_DATABASE are constants used to store the database connection details.
const (
	DB_USERNAME = "DB_USERNAME"
	DB_PASSWORD = "DB_PASSWORD"
	DB_HOST     = "DB_HOST"
	DB_DATABASE = "DB_DATABASE"
)

// DbClient is a global variable that holds the database connection.
var (
	DbClient *sql.DB
)

// init initializes the database connection.
// It loads the environment variables from.env, constructs the data source name, and opens a connection to the database.
// If the connection fails, it panics.
func init() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env.postgres file")
	}

	dbUsername := os.Getenv(DB_USERNAME)
	dbPassword := os.Getenv(DB_PASSWORD)
	dbHost := os.Getenv(DB_HOST)
	dbName := os.Getenv(DB_DATABASE)

	dataSourceName := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", dbUsername, dbPassword, dbHost, dbName)
	var err error
	DbClient, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = DbClient.Ping(); err != nil {
		panic(err)
	}
	log.Println("Connected to DB Server successfully")
}
