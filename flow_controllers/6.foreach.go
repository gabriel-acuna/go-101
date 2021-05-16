package main

import "fmt"

func main() {

	animals := [...]string{"Dog", "Cat", "Fish", "Chicken", "Cow", "Pig"}

	for _, animal := range animals {
		fmt.Println(animal)
	}
}
