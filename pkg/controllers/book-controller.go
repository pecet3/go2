package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error")
	}
	book, err := models.GetBookById(Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if book == nil && err == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

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

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("parsing error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	book, err := models.DeleteBookById(Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error during deleting a book", err)
		return
	}
	if book == nil && err == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var book = &models.Book{}
	fmt.Println(&models.Book{})
	utils.ParseBody(r, book)
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Printf("parsing error")
	}

	bookDetails, err := models.GetBookById(Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if bookDetails == nil && err == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if book.Name == "" {
		book.Name = bookDetails.Name
	}
	if book.Author == "" {
		book.Name = bookDetails.Author
	}
	if book.Publication == "" {
		book.Publication = bookDetails.Publication
	}

	book.Id = Id
	dbBook, err := book.UpdateBookById()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(dbBook)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Conent-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
