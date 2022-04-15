package main

import (
	"fmt"
	"sort"
)

var (
	slice1 = []int{5, 4, 1, 2, 0, -1, 1234, 32, 12, 2345, 99}
)

func main() {
	fmt.Println("Started slice 1 = ", slice1)
	sort.Ints(slice1)
	fmt.Println("Started slice 1 = ", slice1)

	//main1()
}
