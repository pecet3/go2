package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/pecet3/go2/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func initial() {
	fmt.Println("intilal")
	config.Connect()
	fmt.Println("intilal")
	db = config.GetDB()
	if db == nil {
		fmt.Print("errrrrrrr")
	}
	fmt.Println("intilal")
	db.AutoMigrate(&Book{})
	fmt.Println("intilal")
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	// db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)

	return &getBook, db
}

func DeleteBookById(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book
}
