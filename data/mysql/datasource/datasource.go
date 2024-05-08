package datasource

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
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

var (
	DbClient *sql.DB
)

func init() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env.postgres file")
	}

	dbUsername := os.Getenv(DB_USERNAME)
	dbPassword := os.Getenv(DB_PASSWORD)
	dbHost := os.Getenv(DB_HOST)
	dbName := os.Getenv(DB_DATABASE)

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", dbUsername, dbPassword, dbHost, dbName)
	var err error
	DbClient, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = DbClient.Ping(); err != nil {
		panic(err)
	}
	log.Println("Connected to DB Server successfully")
}
