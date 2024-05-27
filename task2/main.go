package main

import (
	"fmt"
	"sync"
)

func square(c int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(c * c)
}

func main() {
	a := [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for _, v := range a {
		wg.Add(1)
		go square(v, &wg)
	}

	wg.Wait()
	fmt.Println("Done")
}
