package routes

import (
	"github.com/gorilla/mux"
	"github.com/pecet3/go2/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/book", controllers.GetBook).Methods("GET")
	router.HandleFunc("/api/book", controllers.CreateBook).Methods("POST")

	router.HandleFunc("/api/book/{bookId}", controllers.GetBookById).Methods("GET")

	router.HandleFunc("/api/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
	// router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
}
