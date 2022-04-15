package main

import (
	"fmt"
	"time"
)

var (
	arr = []int{2, 4, 6, 8, 10}
)

func Square(ch <-chan int) {
	for v := range ch {
		ans := v * v
		fmt.Printf("value %d squared is %d\n", v, ans)
	}
}
func main() {
	ch := make(chan int)
	for i := 0; i < 5; i++ {
		go Square(ch)
		ch <- arr[i]
	}
	time.Sleep(3 * time.Second)

	//main1()
}
