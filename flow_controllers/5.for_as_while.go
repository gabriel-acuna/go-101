package main

import "fmt"

func main() {

	var number, cont int

	fmt.Print("Enter a number:\n")
	fmt.Scanf("%d", &number)

	for number > 0 {
		number = number / 10
		cont++
	}
	fmt.Printf("The number has %d digit(s)", cont)
}
