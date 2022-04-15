package main

import (
	"context"
	"fmt"
	"time"
)

func contextCancel(c context.Context) {
	for {
		select {
		case <-c.Done():
			fmt.Println("goroutine stop by cancel")
			time.Sleep(time.Millisecond * 10)
			return
		}
	}
}

func contextTimeout(c context.Context) {
	for {
		select {
		case <-c.Done():
			fmt.Println("goroutine stop by timeout")
			time.Sleep(time.Millisecond * 10)
			return
		}
	}
}

func chanel(c chan struct{}) {
	for {
		select {
		case <-c:
			fmt.Println("goroutine stop by chanel")
			time.Sleep(time.Millisecond * 10)
			return
		}
	}
}
func main() {
	ctx := context.Background()
	t := time.Second * time.Duration(5)

	ctx1, cancel := context.WithCancel(ctx)
	ctx2, cancel1 := context.WithTimeout(ctx, t)
	ch := make(chan struct{})

	fmt.Println("goroutine start by cancel")
	go contextCancel(ctx1)
	fmt.Println("goroutine start by timeout")
	go contextTimeout(ctx2)
	fmt.Println("goroutine start by chan")
	go chanel(ch)

	select {
	case <-time.After(t):
		cancel()
		cancel1()
		ch <- struct{}{}
	}
	time.Sleep(time.Second * 1)

	//main1()
}
