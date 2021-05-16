package main

import "fmt"

func main() {

	var numbers [5]int
	fruits := [5]string{"orange", "apple", "peach", "lemon", "avocado"}

	numbers[0] = 100
	numbers[1] = 200
	numbers[2] = 300
	numbers[3] = 400
	numbers[4] = 500

	fmt.Println(numbers)
	fmt.Println(numbers[4])
	fmt.Println(fruits)
	fruits0to3 := fruits[0:3]
	fmt.Println(fruits0to3)
}
