package main

import (
	"fmt"
	"sync"
)

var bufferAmount float64 = 0

type BankAccount struct {
	mu      sync.Mutex
	balance float64
}

/*
func (c *BankAccount) deposit(amount float64) {

	c.mu.Lock()
	defer c.mu.Unlock()
	c.balance = c.balance + amount
	fmt.Println("Amount Deposited - ", amount)
}
*/

func doIncrement(amount float64) {
	bufferAmount = bufferAmount + amount
}

func (c *BankAccount) updateBalance() {
	if bufferAmount > 0 {
		c.balance = c.balance + bufferAmount
		fmt.Println("Amount Deposited - ", bufferAmount)
	}
	bufferAmount = 0
}

func (c *BankAccount) withdraw(amount float64) {

	c.mu.Lock()
	defer c.mu.Unlock()
	c.updateBalance()
	if c.balance >= amount {
		c.balance = c.balance - amount
	} else {
		fmt.Println("Withdrawal Failed - Balance Insufficient")
	}
}
func main() {
	c := BankAccount{balance: 0}

	var wg sync.WaitGroup
	/*
		doIncrement := func(n float64) {
			c.deposit(n)
			wg.Done()
		}
	*/
	doWithdrawal := func(n float64) {
		c.withdraw(n)
		wg.Done()
	}

	wg.Add(4)
	go doIncrement(80000)

	go doIncrement(10000)

	go doWithdrawal(1000)

	// go c.updateBalance()

	wg.Wait()
	fmt.Println(c.balance)
}
