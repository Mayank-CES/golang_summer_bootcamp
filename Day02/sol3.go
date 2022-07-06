package main

import (
	"fmt"
	"sync"
)

var amount float64 = 0

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
	} else {
		fmt.Println("Withdrawal Failed - Balance Insufficient")
	}
}
func main() {
	c := BankAccount{balance: 0}

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
	go doIncrement(80000)
	go doIncrement(10000)

	go doWithdrawal(100000)

	wg.Wait()
	fmt.Println(c.balance)
}
