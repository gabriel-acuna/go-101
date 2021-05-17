package main

import "fmt"

func greet(name string){
	fmt.Println("Hi", name)
}

func main(){
    
    var name string
    fmt.Print("What's your name?\n")
    fmt.Scanf("%s", &name)
    greet(name)
}
