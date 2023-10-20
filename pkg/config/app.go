package config

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func initial() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("error")
	}

}

var (
	db *gorm.DB
)

func Connect() {
	dbPassword := os.Getenv("DB_PASSWORD")
	dsn := "user:" + dbPassword + "@tcp(jakubpacewi.cz:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db = d

}

func GetDB() *gorm.DB {
	return db
}
