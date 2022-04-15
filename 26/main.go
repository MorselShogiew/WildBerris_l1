package main

import (
	"fmt"
	"strings"
)

func isUniqMap(a []rune) bool {

	hash := make(map[rune]struct{})
	for i := 0; i < len(a); i++ {
		if _, ok := hash[a[i]]; ok {
			return false
		} else {
			hash[a[i]] = struct{}{}
		}
	}
	return true
}
func main() {
	var (
		slice = []string{
			"AB",
			"世界",
			"世界界",
			" ",
			"  ",
			"❤💛💛💔",
		}
	)
	runes := []rune(strings.ToUpper(slice[0])) // поднимаем все руны в строке вверх
	fmt.Println(isUniqMap(runes))
}
