package main

import "fmt"

func main() {
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September"}
	fmt.Println(months)

	length := len(months)
	capacity := cap(months)

	fmt.Printf("Lenght: %d \nCapacity: %d\n", length, capacity) // max capacity is 9

	months = append(months, "October") // Add a new array
	fmt.Println(months)

	length = len(months)
	capacity = cap(months)

	fmt.Printf("Lenght: %d \nCapacity: %d\n", length, capacity)

}
