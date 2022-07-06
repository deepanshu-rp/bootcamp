package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func addRating(sum *int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	*sum += rand.Intn(10)
	//*sum += 10
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	const students = 200
	sum := 0

	for i:=0; i<students; i++ {
		wg.Add(1)
		go addRating(&sum, &mu, &wg)
	}
	wg.Wait()

	average := sum/students
	fmt.Println(average)
}
