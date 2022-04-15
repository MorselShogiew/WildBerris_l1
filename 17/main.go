package main

import "fmt"

var (
	slice1          = []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	searchedNumber1 = 23
)

func binarySearch(a []int, search int) (result int, searchCount int) {
	mid := len(a) / 2
	switch {
	case len(a) == 0:
		result = -1 // not found
	case a[mid] > search:
		result, searchCount = binarySearch(a[:mid], search)
	case a[mid] < search:
		result, searchCount = binarySearch(a[mid+1:], search)
		if result >= 0 { // if anything but the -1 "not found" result
			result += mid + 1
		}
	default: // a[mid] == search
		result = mid // found
	}
	searchCount++
	return
}
func main() {
	index, c := binarySearch(slice1, searchedNumber1)
	fmt.Println(index, c)

	//main1()
}

//берем отсортированный массив, смотрим в середину, если не нашли там число,
//в зависимости от того, что в середине — ищем это число этим же методом либо в левой части, либо в правой,
//откидывая средний элемент. Для функций также, просто берем не массив, а функцию.
