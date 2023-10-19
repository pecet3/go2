package main

import (
	"fmt"
)

type User struct {
	name string
	id   int
}

func (u User) sayName() {
	fmt.Printf("hello %s\n", u.name)
}

func main() {

	user := User{
		name: "jakub",
		id:   1,
	}

	User.sayName(user)
}
