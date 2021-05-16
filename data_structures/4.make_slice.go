package main

import "fmt"

func main() {

	scores := make([]float32, 3, 3)

	scores[0] = 7.9
	scores[1] = 8.97
	scores[2] = 10

	scores = append(scores, 9.79)

	fmt.Println(scores)
	fmt.Println(len(scores))
	fmt.Println(cap(scores))
}
