package main

import "fmt"

func main() {
	var score float32

	fmt.Println("Enter the score:")
	fmt.Scanf("%f", &score)
	switch {
	case score == 10:
		fmt.Println("Excellent!")
		break
	case score == 9 || score == 8:
		fmt.Println("Very Good!")
		break
	case score == 7 || score == 6:
		fmt.Println("Not bad, but try harder next time!")
		break
	case score >= 0 && score <= 5:
		fmt.Println("Sorry you didn't pass!")
		break
	default:
		fmt.Println("Score is not valid!")

	}
}
