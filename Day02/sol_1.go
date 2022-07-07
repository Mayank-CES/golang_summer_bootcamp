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

func printFrequency(result [26]int) {
	for i, count := range result {
		if count > 0 {
			fmt.Printf(" %q:%d", 97+i, count)
		}
	}

}

func main() {

	ch := make(chan string, 5)
	var dataset = [5]string{"quick", "brown", "fox", "lazy", "dog"}
	var result = [26]int{}

	wg.Add(5)

	go getFrequency(ch, &result)

	for _, word := range dataset {
		ch <- word
	}
	defer close(ch)

	wg.Wait()

	fmt.Printf("{")
	printFrequency(result)
	fmt.Printf(" }")
}
