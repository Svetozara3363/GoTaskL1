package main

import (
	"fmt"
	"sync"
)

var (
	sum int
	mu  sync.Mutex
	wg  sync.WaitGroup
)

func square(k int) {
	defer wg.Done()
	result := k * k
	mu.Lock()
	sum += result
	mu.Unlock()
}

func main() {
	a := [5]int{2, 4, 6, 8, 10}
	for _, val := range a {
		wg.Add(1)
		go square(val)
	}
	wg.Wait()
	fmt.Println("Sum is -", sum)
}
