package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pecet3/go2/pkg/models"
	"github.com/pecet3/go2/pkg/utils"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {

	newBooks, err := models.GetAllBooks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error during cfetching the books")
		return
	}
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

// func GetBookById(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	bookId := vars["bookId"]

// 	Id, err := strconv.ParseInt(bookId, 0, 0)
// 	if err != nil {
// 		fmt.Println("error")
// 	}
// 	bookDetails, _ := models.GetBookById(Id)
// 	res, _ := json.Marshal(bookDetails)

// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)

// }

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)
	b, err := book.CreateBook()
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error during creating a book")
		return
	}
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// func DeleteBook(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	bookId := vars["bookId"]
// 	Id, err := strconv.ParseInt(bookId, 0, 0)
// 	if err != nil {
// 		fmt.Printf("parsing error")
// 	}

// 	book := models.DeleteBookById(Id)
// 	res, _ := json.Marshal(book)

// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// func UpdateBook(w http.ResponseWriter, r *http.Request) {
// 	var book = &models.Book{}
// 	utils.ParseBody(r, book)
// 	vars := mux.Vars(r)
// 	bookId := vars["bookId"]

// 	Id, err := strconv.ParseInt(bookId, 0, 0)
// 	if err != nil {
// 		fmt.Printf("parsing error")
// 	}

// 	bookDetails, db := models.GetBookById(Id)

// 	if book.Name != "" {
// 		bookDetails.Name = book.Name
// 	}
// 	if book.Author != "" {
// 		bookDetails.Author = book.Name
// 	}
// 	if book.Publication != "" {
// 		bookDetails.Publication = book.Publication
// 	}

// 	res, _ := json.Marshal(bookDetails)
// 	w.Header().Set("Conent-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }
