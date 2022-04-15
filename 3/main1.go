package main

import (
	"fmt"
)

var (
	arr = []int{2, 4, 6, 8, 10}
)

func Square(ch1 <-chan int, ch2 chan<- int) {
	var ans int
	for v := range ch1 {
		ans += v * v
	}
	ch2 <- ans
}
func main1() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Square(ch1, ch2)
	for _, v := range arr {
		ch1 <- v
	}
	close(ch1)
	ans := <-ch2
	fmt.Println("answer is", ans)
}
