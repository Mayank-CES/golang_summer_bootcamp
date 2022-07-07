package main

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	mu      sync.Mutex
	balance float64
}

func (c *BankAccount) deposit(amount float64, wg *sync.WaitGroup) {
	defer wg.Done()
	c.mu.Lock()
	defer c.mu.Unlock()

	c.balance = c.balance + amount
	fmt.Println("Amount Deposited - ", amount)
}

func (c *BankAccount) withdraw(amount float64, wg *sync.WaitGroup) {
	defer wg.Done()
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

	wg.Add(3)

	go c.deposit(800, &wg)
	go c.deposit(100, &wg)

	go c.withdraw(200, &wg)

	wg.Wait()
	fmt.Println(c.balance)
}
