package main

import "fmt"

var (
	slice21 = []int{1, 4, 3, 2, 8}
	slice22 = []int{5, 1, 7, 4, 8, 9}
)

func intersection(s1, s2 []int) (inter []int) {
	hash := make(map[int]bool)
	for _, e := range s1 {
		hash[e] = true
	}
	for _, e := range s2 {
		if hash[e] {
			inter = append(inter, e)
		}
	}
	inter = removeDups(inter)
	return
}

func removeDups(elements []int) (nodups []int) {
	encountered := make(map[int]bool)
	for _, element := range elements {
		if !encountered[element] {
			nodups = append(nodups, element)
			encountered[element] = true
		}
	}
	return
}
func main() {
	slice23 := []int{}
	slice23 = intersection(slice21, slice22)
	fmt.Println(slice23)

	//main1()
}
