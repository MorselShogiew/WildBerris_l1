package main

import (
	"fmt"
	"strings"
)

func main1() {
	text := "snow dog sun"
	split := strings.Split(text, " ")
	var result []string
	for i := len(split) - 1; i >= 0; i-- {
		result = append(result, split[i])
	}
	fmt.Println(result)
}
