package main

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	mu      sync.Mutex
	balance float64
}

func (c *BankAccount) deposit(amount float64) {

	c.mu.Lock()
	defer c.mu.Unlock()
	c.balance = c.balance + amount
	fmt.Println("Amount Deposited - ", amount)
}

func (c *BankAccount) withdraw(amount float64) {

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.balance >= amount {
		c.balance = c.balance - amount
		fmt.Println("Amount Withdrawn - ", amount)
	} else {
		fmt.Println("Withdrawal Failed - Balance Insufficient")
	}
}
func main() {
	c := BankAccount{balance: 500}

	var wg sync.WaitGroup

	doIncrement := func(n float64) {
		c.deposit(n)
		wg.Done()
	}

	doWithdrawal := func(n float64) {
		c.withdraw(n)
		wg.Done()
	}

	wg.Add(3)
	go doIncrement(800)
	go doIncrement(100)

	go doWithdrawal(200)

	wg.Wait()
	fmt.Println(c.balance)
}
