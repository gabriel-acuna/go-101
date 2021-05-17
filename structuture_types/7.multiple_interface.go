package main

import "fmt"

type Animal interface {
	Sleep()
}

type Pet interface {
	Eat()
}

type Cat struct {
	Name string
}

func (self Cat) Sleep() {
	fmt.Printf("%s is sleepping\n", self.Name)
}

func (self Cat) Eat() {
	fmt.Printf("%s is eatting\n", self.Name)
}

func animalAction(animal Animal) {
	fmt.Print("This object is an animal!")
}

func petAction(pet Pet) {
	fmt.Print("This object is a pet!")
}
func main() {
	myCat := Cat{Name: "Copy"}

	myCat.Eat()
	myCat.Sleep()
	animalAction(myCat)
	petAction(myCat)

}
