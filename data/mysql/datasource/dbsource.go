package datasource

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

const (
	DB_USERNAME = "DB_USERNAME"
	DB_PASSWORD = "DB_PASSWORD"
	DB_HOST     = "DB_HOST"
	DB_PORT     = "DB_PORT"
	DB_DATABASE = "DB_DATABASE"
)

var DbClient *sql.DB

func init() {

	if err := godotenv.Load(".env.postgres"); err != nil {
		log.Fatal("Error loading .env.postgres file")
	}

	dbUsername := os.Getenv(DB_USERNAME)
	dbPassword := os.Getenv(DB_PASSWORD)
	dbHost := os.Getenv(DB_HOST)
	dbPort := os.Getenv(DB_PORT)
	dbName := os.Getenv(DB_DATABASE)

	dataSourceName := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		dbUsername,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)
	DbClient, err := sql.Open(
		"postgres",
		dataSourceName,
	)
	if err != nil {
		panic(err)
	}

	if err = DbClient.Ping(); err != nil {
		panic(err)
	}
	log.Println("Connected to DB Server successfully")
}
