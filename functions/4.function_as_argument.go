package main

import "fmt"

type Transaction func(balance, amount int) int

func desposit(balance, amount int) int {
	return balance + amount
}

func withdraw(balance, amount int) int {
	return balance - amount
}

func executeTransaction(function Transaction, balance, amount int) {
	fmt.Println("Before the transaction")
	result := function(balance, amount)
	fmt.Println("After the transaction")
	fmt.Println("Your current balance is:", result)

}

func main() {
	executeTransaction(desposit, 100, 80)
	executeTransaction(withdraw, 300, 75)
}
