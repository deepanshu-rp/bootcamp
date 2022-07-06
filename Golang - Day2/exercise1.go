package main

import (
	"fmt"
	"sync"
)

// count letters in word
func count(ch <-chan string, result *[]int, wg *sync.WaitGroup) {
	for word := range ch {
		for _, char := range word {
			(*result)[char-97] += 1
		}
		wg.Done()
	}
}

// print as char:count
func printResult(result []int) {
	ans := make(map[string]int)
	for i:= range result {
		ans[string(97+i)] = result[i]
	}
	fmt.Println(ans)
}

func main() {
	var words = []string{"quick", "brown", "fox", "lazy", "dog"}
	var result = new([26]int)[0:26]
	var wg sync.WaitGroup

	// channel with buffer size of number of words
	ch := make(chan string, len(words))

	wg.Add(len(words))
	go count(ch, &result, &wg)

	// add words to channel
	for _, word := range words {
		ch <- word
	}
	defer close(ch)
	wg.Wait()

	printResult(result)
}
