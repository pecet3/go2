package config

import (
	"database/sql"
	"os"

	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
)

var (
	db *sql.DB
)

func Connect() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("error")
	}

	pass := os.Getenv("DB_PASSWORD")

	d, err := sql.Open("mysql", "user:"+pass+"@tcp(jakubpacewi.cz)/db1")
	if err != nil {
		panic(err.Error())
	}

	err = d.Ping()
	if err != nil {
		panic(err.Error())
	}

	db = d

}

func GetDB() *sql.DB {
	return db
}
