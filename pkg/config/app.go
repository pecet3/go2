package config

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
	dbPassword := ""
	dsn := "user:" + dbPassword + "@tcp(jakubpacewi.cz:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Error with db")
	}
	fmt.Println(d)
	db = d

}

func GetDB() *gorm.DB {
	fmt.Println(db)
	return db
}
