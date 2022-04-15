package main

import "fmt"

var (
	slice1 = []int{1, 4, 3, 2, 8}
	slice2 = []int{5, 1, 7, 4, 8, 9}
)

func TotalS(s1 []int, s2 []int) []int {
	var cross []int
	for _, val1 := range s1 {
		for _, val2 := range s2 {
			if val1 == val2 {
				cross = append(cross, val1)
			}
		}
	}
	return cross
}

func main1() {
	cross := TotalS(slice1, slice2)
	fmt.Println("Crossing: ", cross)

	//main1()
}
