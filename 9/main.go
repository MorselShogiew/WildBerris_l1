package main

import (
	"fmt"
)

var arr = []int{2, 4, 6, 8, 10}

func Source(downstream chan int) {
	for _, v := range arr {
		downstream <- v
	}
	close(downstream)
}
func Filter(upstream, downstream chan int) {
	for item := range upstream {
		downstream <- item * item
	}
	close(downstream)
}
func Print(upstream chan int) {
	for v := range upstream {
		fmt.Printf("value squared is %d\n", v)
	}

}
func main() {
	c0 := make(chan int)
	c1 := make(chan int)
	go Source(c0)
	go Filter(c0, c1)
	Print(c1)
}
