// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// 	"github.com/pecet3/go2/pkg/routes"
// )

// func main() {
// 	fmt.Printf("Listening")
// 	r := mux.NewRouter()
// 	routes.RegisterBookStoreRoutes(r)
// 	fmt.Printf("Listening 2")
// 	http.Handle("/", r)
// 	fmt.Printf("Listening 3")

// 	if err := http.ListenAndServe("localhost:8080", r); err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("Listening 4")

// }

package main

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

func main() {
	dbPassword := "Reksiokuba1!"
	dsn := "user:" + dbPassword + "@tcp(jakubpacewi.cz:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Error with db")
	}
	fmt.Println(d)
	db = d
	fmt.Println(db)
}

func GetDB() *gorm.DB {
	fmt.Println(db)
	return db
}
