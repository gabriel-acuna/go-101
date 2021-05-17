package main

import "fmt"

func factorial(number int) int {
	if number >= 0 && number < 2 {
		return 1
	}

	return factorial(number-1) * number

}

func main() {
	var number int
	fmt.Print("Enter a number:\n")
	fmt.Scanf("%d", &number)
	result := factorial(number)
	fmt.Printf("%d! = %d\n", number, result)
}
