package main

import "fmt"

var (
	slice3 = []int{0, 1, 2, 3, 4, 5, 6, 7}
	index3 = 3
)

func main2() {
	slice4 := append((slice3)[0:index3], slice3[index3+1:]...)
	fmt.Println(slice4)
}
