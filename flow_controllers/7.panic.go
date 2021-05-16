package main

import "fmt"

func main() {

	var dividend, divider float32

	fmt.Print("Enter the dividend:")
	fmt.Scanf("%f", &dividend)

	fmt.Print("Enter the divider:")

	fmt.Scanf("%f", &divider)

	if divider == 0 {
		panic("The divider should be different than zero")
	} else {
		fmt.Printf("Result: %f", dividend/divider)
	}

}
