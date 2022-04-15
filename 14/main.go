package main

import (
	"fmt"
)

var (
	a interface{}
)

func main() {
	func(variable interface{}) {
		switch variable.(type) {
		case int:
			fmt.Printf("%v is int\n", variable)
		case string:
			fmt.Printf("%v is string\n", variable)
		case bool:
			fmt.Printf("%v is bool\n", variable)
		case chan int:
			fmt.Printf("%v is chan int\n", variable)
		default:
			fmt.Println("Unknown type")
		}
	}(a)

}
