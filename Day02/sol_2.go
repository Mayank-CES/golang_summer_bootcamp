package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var Score float64 = 0

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	rand.Seed(time.Now().UnixNano())
	Score = Score + float64(rand.Intn(10)) + 1
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 200; i++ {
		wg.Add(1)

		i := i

		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()

	fmt.Println("Score : ", (Score / 200.0))
}
