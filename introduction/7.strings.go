package main

import (
	"fmt"
	"reflect"
)

func main() {
	var course string = "Go Professional Course"
	var instructor = "Eduardo García"
	platform := "Código Facilito"
	fmt.Println("Course:", course)
	fmt.Println("Instructor:", instructor)
	fmt.Println("Platform:")

	// obtner el tamamño del string
	fmt.Println(len(course))

	firstCharacter := platform[0]

	fmt.Println(firstCharacter)

	fmt.Println(reflect.TypeOf(firstCharacter))

	fmt.Printf("%c\n", firstCharacter)
}
