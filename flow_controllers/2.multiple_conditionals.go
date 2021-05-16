package main

import "fmt"

func main() {

	var score float32

	fmt.Println("Enter the score:")
	fmt.Scanf("%f", &score)
	if score == 10 {
		fmt.Println("Excellent!")
	} else if score < 10 && score > 7 {
		fmt.Println("Very Good!")
	} else if score < 8 && score > 5 {
		fmt.Println("Not bad, but try harder next time!")
	} else if score >= 0 && score <= 5 {
		fmt.Println("Sorry you didn't pass!")
	} else {
		fmt.Println("Score is not valid!")
	}
}
