package main

import "fmt"

func function1() {
	fmt.Println("Hi from Function1!")
}

func function2() {
	fmt.Println("Hi from Function2!")
}

func main() {

	defer function1()
	function2()
}
