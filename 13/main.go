package main

import "fmt"

func main() {
	a := 15
	b := 13
	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)
	a = b + a
	b = a - b
	a = a - b
	//a = a * b; b = a / b; a = a / b;
	fmt.Println(a, b)
}
