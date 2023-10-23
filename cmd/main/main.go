package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pecet3/go2/pkg/routes"
)

func main() {

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	http.Handle("/api", r)

	port := 8080

	fmt.Printf("Starting server at port:%v\n", port)

	servePort := ":" + strconv.Itoa(port)

	if err := http.ListenAndServe(servePort, r); err != nil {
		log.Fatal(err)
	}
}
