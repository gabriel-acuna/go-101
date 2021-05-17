package main

import (
	"fmt"

	"./MyPackage"
)

func main() {
	person := MyPackage.Person{FirstName: "Gabriel", LastName: "Acuña", Height: 1.68, Weight: 81.4}
	fmt.Println(person.GetFullName())

}
