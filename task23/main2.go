package main

import "fmt"

func main() {
	a := []string{"A", "B", "C", "D", "E"}
	i := 2
	copy(a[i:], a[i+1:])

	a = a[:len(a)-1]

	fmt.Println(a)

	// done A B D E
}
