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
	DB_DATABASE = "DB_DATABASE"
)

var DbClient *sql.DB

func init() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUsername := os.Getenv(DB_USERNAME)
	dbPassword := os.Getenv(DB_PASSWORD)
	dbHost := os.Getenv(DB_HOST)
	dbName := os.Getenv(DB_DATABASE)

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", dbUsername, dbPassword, dbHost, dbName)
	DbClient, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = DbClient.Ping(); err != nil {
		panic(err)
	}
	//mysql.SetLogger()
	log.Println("Connected to MySQL DB Server successfully")
}
