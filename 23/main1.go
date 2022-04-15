package main

import "fmt"

var (
	slice2 = []int{0, 1, 2, 3, 4, 5, 6, 7}
	index2 = 3
)

func main1() {
	fmt.Println("Started slice = ", slice2)
	slice2[index2] = slice2[len(slice2)-1]
	slice2 = slice2[:len(slice2)-1]
	fmt.Println("Finished slice = ", slice2)
}
