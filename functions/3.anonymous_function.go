package main

import "fmt"

func main() {
	sayHiWorld := func(name string) string {
		message := fmt.Sprintf("Hi %s!", name)
		return message
	}

	message := sayHiWorld("Gabo")
	fmt.Println(message)
}
