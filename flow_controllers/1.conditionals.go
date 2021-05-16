package main

import "fmt"

func main() {
	var number int
	fmt.Println("Enter a number:")
	fmt.Scanf("%d", &number)
	if number > 10 {
		fmt.Println("the number that you entered is greater or equal than 10")
	} else {
		fmt.Println("the number that you entered is lower than 10")
	}

}
