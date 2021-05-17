package main

import "fmt"

func average(scores ...float32) float32 {
	var total float32
	for _, score := range scores {
		total = total + score
	}
	return total / float32(len(scores))
}

func main() {
	finalScore := average(10, 8, 8.5, 9, 9.88)
	fmt.Println(finalScore)
}
