package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Active bool
}

type Student struct {
	User User
	Id   string
}

func main() {

	user1 := User{Name: "Gabriel", Email: "g.acuna@mail.com", Active: true}

	user2 := User{Name: "Stefano", Email: "s.acuna93@mail.com", Active: true}
	fmt.Println(user2)

	student := Student{User: user1, Id: "A09GY"}
	fmt.Println(student)
}
