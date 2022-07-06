package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func getFrequency(ch <-chan string, result *[26]int) {
	for word := range ch {
		for _, char := range word {
			(*result)[char-97] += 1
		}
		wg.Done()
	}
}

func main() {
	// pings := make(chan string, 1)
	// pongs := make(chan string, 1)

	ch := make(chan string, 5)
	var dataset = [5]string{"quick", "brown", "fox", "lazy", "dog"}
	var result = [26]int{26 * 0}

	wg.Add(5)

	go getFrequency(ch, &result)

	for _, word := range dataset {
		ch <- word
	}
	defer close(ch)

	wg.Wait()

	fmt.Println(result)

}
