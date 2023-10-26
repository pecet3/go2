package models

import (
	"database/sql"
	"fmt"

	"github.com/pecet3/go2/pkg/config"
)

var db *sql.DB

type Book struct {
	Id          int64  `json:"ID"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func initial() {

	config.Connect()

	db = config.GetDB()

	if db == nil {
		fmt.Print("error during connect db")
	}

}

func (b *Book) CreateBook() (*Book, error) {
	initial()
	config.Connect()

	_, err := db.Exec("INSERT INTO books (name, author, publication) VALUES (?, ?, ?)", b.Name, b.Author, b.Publication)
	if err != nil {
		return nil, err
	}

	return b, nil

}
func GetAllBooks() ([]Book, error) {
	initial()
	var Books []Book
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book Book
		err = rows.Scan(&book.Id, &book.Name, &book.Author, &book.Publication)
		if err != nil {
			panic(err.Error())
		}

		Books = append(Books, book)

		fmt.Printf("ID: %d, Name: %s, Author: %s, Publication: %s\n", book.Id, book.Name, book.Author, book.Publication)
	}

	if err := rows.Err(); err != nil {
		panic(err.Error())
	}
	return Books, nil
}

func GetBookById(Id int64) (*Book, error) {
	isFound := false
	books, err := GetAllBooks()
	if err != nil {
		return nil, err
	}

	var book Book

	for b := range books {

		if books[b].Id == Id {
			book = books[b]
			isFound = true
			break
		}

	}
	if isFound == false {
		return nil, err
	}

	return &book, nil
}

func DeleteBookById(Id int64) (*Book, error) {
	initial()

	var book Book
	row, err := db.Query("SELECT * FROM books WHERE ID = ?", Id)

	for row.Next() {
		err = row.Scan(&book.Id, &book.Name, &book.Author, &book.Publication)
		if err != nil {
			fmt.Println("error scaning")
			return nil, err
		}
	}
	if err := row.Err(); err != nil {
		panic(err.Error())
	}

	responseDb, err := db.Exec("DELETE FROM books WHERE ID = ?", Id)
	if err != nil {
		return nil, err
	}
	fmt.Println(responseDb)
	return &book, nil

}
