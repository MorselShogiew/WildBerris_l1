package main

import (
	"fmt"
	"sort"
)

var (
	slice21          = []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	searchedNumber12 = 38
)

func main1() {
	fmt.Println("Slice = ", slice1)
	i := sort.Search(len(slice1), func(i int) bool { return slice1[i] >= searchedNumber1 })
	if i < len(slice1) && slice1[i] == searchedNumber1 {
		fmt.Printf("%v is index of %v", searchedNumber1, i)
	} else {
		fmt.Printf("%d not found in %v\n", searchedNumber1, slice1)
	}
}
