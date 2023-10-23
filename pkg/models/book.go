package models

import (
	"database/sql"
	"fmt"

	"github.com/pecet3/go2/pkg/config"
)

var db *sql.DB

type Book struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func Initial() {

	config.Connect()
	
}
	db = config.GetDB()

	if db == nil {
		fmt.Print("errrrrrrr")
	}

	

func (b *Book) CreateBook() *Book {

	return b
}
func GetAllBooks() []Book {
	config.Connect()
	defer db.Close()
	var Books []Book
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var book Book
		var id int
		err = rows.Scan(&id, &book.Name, &book.Author, &book.Publication)
		if err != nil {
			panic(err.Error())
		}

		Books = append(Books, book)

		fmt.Printf("ID: %d, Name: %s, Author: %s, Publication: %s\n", book.Name, book.Author, book.Publication)
	}

	if err := rows.Err(); err != nil {
		panic(err.Error())
	}
	return Books
}

func GetBookById(Id int64) (*Book, *sql.DB) {
	var getBook Book

	return &getBook, db
}

func DeleteBookById(Id int64) Book {
	var book Book

	return book
}
