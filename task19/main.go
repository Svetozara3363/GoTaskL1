package main

import (
	"fmt"
)

func reverse(str string) string {
	sliceRune := []rune(str)
	for i, j := 0, len(sliceRune)-1; i < j; i, j = i+1, j-1 {
		sliceRune[i], sliceRune[j] = sliceRune[j], sliceRune[i]
	}
	return string(sliceRune)
}

func main() {
	str := "string"
	newString := reverse(str)
	fmt.Println(newString)
}
