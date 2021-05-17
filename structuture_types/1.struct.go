package main

import "fmt"

type User struct {
	Name  string
	Email string
}

func main() {
	var user1 User
	user1.Name = "Gabriel"
	user1.Email = "g.acuna@mail.com"
	fmt.Println(user1)

	user2 := User{"Jesenia", "jessy.repi@mail.com"}
	fmt.Println(user2)

	Name := "Stefano-Acuña"
	Email := "s.acuna93@mail.com"

	user3 := User{Name, Email}

	fmt.Println(user3)
}
