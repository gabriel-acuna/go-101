package main

import "fmt"

type UserType int

const (
	Teacher UserType = 1
	Student UserType = 2
)

type User struct {
	Username string
	UserType UserType
}

func main() {
	student := User{Username: "Gabriel", UserType: Student}
	fmt.Println(student)

	teacher := User{Username: "Stefano", UserType: Teacher}
	fmt.Println(teacher)
}
