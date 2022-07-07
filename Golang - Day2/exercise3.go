package main

import (
	"fmt"
	"sync"
)


func deposit(balance *int, amount int, mu *sync.Mutex, wg *sync.WaitGroup) {
	mu.Lock()
	fmt.Println(*balance, amount, *balance+amount)
	*balance += amount
	fmt.Println("Previous:", *balance-100, "Current:", *balance)
	mu.Unlock()

	wg.Done()
}


func withdraw(balance *int, amount int, mu *sync.Mutex, wg *sync.WaitGroup) {
	mu.Lock()
	// update balance if transaction valid else throw error
	if v := *balance-amount; v>= 0 {
		fmt.Println(*balance, amount, *balance-amount)
		*balance = v
		fmt.Println("Previous:", *balance+100, "Current:", *balance)

	} else {
		panic(any("Invalid transaction"))
	}
	mu.Unlock()

	wg.Done()
}

func main() {
	var balance = 500
	var wg sync.WaitGroup
	var mu sync.Mutex

	const iterations = 6
	wg.Add(2*iterations)

	for i:=0; i<iterations; i++ {
		go withdraw(&balance,100, &mu, &wg)
		go deposit(&balance,100, &mu, &wg)
	}

	wg.Wait()

	fmt.Println(balance)

}
