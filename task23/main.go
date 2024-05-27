package main

import "fmt"

func main() {
	a := []string{"A", "B", "C", "D", "E"}
	i := 2
	a[i] = a[len(a)-1]

	a = a[:len(a)-1]

	fmt.Println(a)
}
