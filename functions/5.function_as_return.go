package main

import "fmt"

type Transaction func(balance, quantity float32) float32

func createTrasaction(transactionType string) Transaction {

	if transactionType == "Deposit" {
		return func(balance, amount float32) float32 { return balance + amount }
	} else if transactionType == "Withdraw" {
		return func(balance, amount float32) float32 { return balance - amount }
	} else {
		return func(balance, amount float32) float32 { return balance + amount*0.05 }
	}

}

func main() {
	desposit := createTrasaction("Deposit")
	currentBalance := desposit(1000, 679.76)
	fmt.Printf("Your current balance is: %.2f\n", currentBalance)

}
