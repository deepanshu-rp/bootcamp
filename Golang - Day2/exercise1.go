package main

import (
	"fmt"
	"sync"
)

func count(ch <-chan string, result *[]int, mu *sync.Mutex) {
	for word := range ch {
		for _, char := range word {
			//mu.Lock()
			(*result)[char-97] += 1
			//mu.Unlock()
		}
	}
}

func main() {
	//var result map[string]int
	var result = new([26]int)[0:26]
	//const workers = 3
	var mu sync.Mutex
	ch := make(chan string, 5)
	var words = []string{"quick", "brown", "fox", "lazy", "dog"}
	go count(ch, &result, &mu)
	//for j:=0; j<workers; j++ {
	//	go count(ch, &result)
	//}

	for _, word := range words {
		ch <- word
	}
	defer close(ch)

	fmt.Println(result)
	//ans := make(map[string]int)
	//for i:= range result {
	//	ans[string(97+i)] = result[i]
	//}
	//fmt.Println(ans)

}
