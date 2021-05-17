package main

import (
	"errors"
	"fmt"
)

func divide(dividend, divider float32) (float32, error) {
	if divider == 0 {
		return 0, errors.New("It is not possible to divide by zero")
	} else {
		return dividend / divider, nil
	}
}

func main() {

	var dividend, divider float32

	fmt.Print("Enter the dividend: ")
	fmt.Scanf("%f", &dividend)

	fmt.Print("Enter the divider: ")
	fmt.Scanf("%f", &divider)

	if result, err := divide(dividend, divider); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The result is:", result)
	}
}
