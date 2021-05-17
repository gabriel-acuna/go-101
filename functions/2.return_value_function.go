package main

import "fmt"

func add( num1 int, num2 int) int{
	return num1 + num2
}

func main(){
	result := add(2,8)
	fmt.Println(result)
}