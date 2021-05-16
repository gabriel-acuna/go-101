package main

import (
	"fmt"
	"reflect"
)

func main() {
	numbers := []int{1, 2, 3, 4}
	fmt.Println(reflect.TypeOf(numbers))

	fmt.Println(numbers)
	numbers = append(numbers, 5)
	numbers = append(numbers, 6)
	fmt.Println(numbers)
}
