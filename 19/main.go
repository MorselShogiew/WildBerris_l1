package main

import (
	"fmt"
)

var s = "世界, string, ⌘, Марсель"

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println("Before ", s)

	rev := reverse(s)

	fmt.Println("After ", rev)
}
