package main

import "fmt"

type BankAccount struct {
	owner  string
	balance float64
}

// deposite method to add money to the account
func (account *BankAccount) Deposite(amount float64) {
	account.balance += amount
	fmt.Println("Deposited: ", amount, "New Balance: ", account.balance)
}

// withdraw method to remove money from the account
func (account *BankAccount) Withdraw(amount float64) {
	if account.balance >= amount {
		account.balance -=amount
		fmt.Println("Withdrawn: ", amount, "New Balance: ", account.balance)
	} else {
		fmt.Println("Insufficient funds. Current balance: ", account.balance)
	}
}

// check balance method to view the current balance
func (account BankAccount) CheckBalance() {
	fmt.Println("Current Balance for ",account.owner, "is : ", account.balance)
}

func main() {
	// Create a new bank account
	account := BankAccount{owner: "John Doe", balance: 1000.0}

	// Check the initial balance
	account.CheckBalance()

	// Deposit some money
	account.Deposite(500.0)

	// Withdraw some money
	account.Withdraw(200.0)

	// Check the final balance
	account.CheckBalance()
}