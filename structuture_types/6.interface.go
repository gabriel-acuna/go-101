package main

import "fmt"

type Animal interface {
	Eat()
	Sleep()
}

type Dog struct {
	Name string
}

func (self Dog) Eat() {
	fmt.Printf("%s is eatting \n", self.Name)
}
func (self Dog) Sleep() {
	fmt.Printf("%s is sleepping \n", self.Name)
}

func executeActions(animal Animal) {
	animal.Eat()
	animal.Sleep()

}
func main() {

	myDog := Dog{Name: "Docky"}
	executeActions(myDog)

}
