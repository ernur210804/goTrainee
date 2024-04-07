package main

import (
	"errors"
	"fmt"
	"sync"
)

// Bitcoin represents the amount of bitcoins in the wallet.
type Bitcoin float64

// Wallet represents a Bitcoin wallet.
type Wallet struct {
	balance Bitcoin
	mutex   sync.Mutex
}

// Deposit adds funds to the wallet.
func (w *Wallet) Deposit(amount Bitcoin) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	w.balance += amount
}

// Withdraw subtracts funds from the wallet.
func (w *Wallet) Withdraw(amount Bitcoin) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	if amount > w.balance {
		return errors.New("insufficient funds")
	}
	w.balance -= amount
	return nil
}

// Balance returns the current balance of the wallet.
func (w *Wallet) Balance() Bitcoin {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	return w.balance
}

func main() {
	var choice int
	var amount Bitcoin
	wallet := Wallet{}

	for {
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Check Balance")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter amount to deposit: ")
			fmt.Scanln(&amount)
			wallet.Deposit(amount)
			fmt.Println("Amount deposited successfully.")
		case 2:
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scanln(&amount)
			err := wallet.Withdraw(amount)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Amount withdrawn successfully.")
			}
		case 3:
			fmt.Println("Current balance:", wallet.Balance())
		case 4:
			fmt.Println("Exiting program.")
			return
		default:
			fmt.Println("Invalid choice. Please enter a valid option.")
		}
		fmt.Println()
	}
}
