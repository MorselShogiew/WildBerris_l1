package main

import (
	"fmt"
	"time"
)

func read(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}
func write(ch chan<- int) {
	var i int
	for {
		ch <- i
		time.Sleep(time.Millisecond * 50)
		i++
	}
}

func main() {
	c := make(chan int)
	var N int
	fmt.Println("write N")
	fmt.Scan(&N)

	go read(c)
	go write(c)
	time.Sleep(time.Duration(N) * time.Second)
	close(c)
	fmt.Println("chanal close")
}
