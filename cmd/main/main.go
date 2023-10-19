package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Println("error")
	}
	b := os.Getenv("PASSWORD")
	fmt.Print(b)
}
