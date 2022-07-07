package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func student(id int, Score *float64, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Printf("Student %d starting\n", id)
	rand.Seed(time.Now().UnixNano())

	*Score = *Score + float64(rand.Intn(10)) + 1

	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Printf("Student %d done\n", id)
}

func main() {

	var Score float64 = 0

	var wg sync.WaitGroup

	for i := 1; i <= 200; i++ {
		wg.Add(1)

		i := i

		go student(i, &Score, &wg)
	}

	wg.Wait()

	fmt.Println("Average Score by 200 Students : ", (Score / 200.0))
}
