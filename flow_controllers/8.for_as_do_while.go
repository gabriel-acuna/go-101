package main

import "fmt"

func main() {

	var number = 100
	for ok := true; ok; ok = number <= 150 {
		fmt.Print(number, " ")
		number++
	}
}
