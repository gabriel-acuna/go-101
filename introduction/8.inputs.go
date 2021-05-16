package main

import "fmt"

func main() {

	var name string
	var age int
	var height float32

	fmt.Print("Enter your name:\n")
	fmt.Scanf("%s", &name)

	fmt.Print("Enter your age:\n")
	fmt.Scanf("%d", &age)

	fmt.Print("Enter your height:\n")
	fmt.Scanf("%f", &height)

	fmt.Printf("Your name is %s\nYou're %d years old \nYour height is %.2f\n", name, age, height)

}
