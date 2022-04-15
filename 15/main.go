package main

import (
	"fmt"
)

func createHugeString(size int) string {
	var v string
	for i := 0; i < size; i++ {
		v += "a"
	}
	return v
}

func someFunc() {
	v := createHugeString(1 << 10)
	res := make([]rune, 0, 100)
	res = []rune(v)[:100]
	justString := string(res)
	//justString := v[:100]

	fmt.Println("v = ", v)
	fmt.Println("justString = ", justString)
	fmt.Println("len(v) = ", len(v))
	fmt.Println("len(justString) = ", len(justString))
	fmt.Println("cap(v) = ", cap([]byte(v)))
	fmt.Println("cap(justString) = ", cap([]byte(justString)))
}

func main() {
	someFunc()
	v := make([]int, 0, 1024)
	for i := 0; i < 1024; i++ {
		v = append(v, i)
	}
	justString := v[:100]
	fmt.Println("v = ", v)
	fmt.Println("justString = ", justString)
	fmt.Println("len(v) = ", len(v))
	fmt.Println("len(justString) = ", len(justString))
	fmt.Println("cap(v) = ", cap(v))
	fmt.Println("cap(justString) = ", cap(justString))
}
