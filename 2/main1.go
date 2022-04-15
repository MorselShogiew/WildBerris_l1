package main

import (
	"fmt"
	"sync"
)

func main1() {
	var wg sync.WaitGroup
	arr := []int{2, 4, 6, 8, 10}
	for _, i := range arr {
		go func(wg *sync.WaitGroup, i int) {
			wg.Add(1)
			fmt.Println(i * i)
			defer wg.Done()
		}(&wg, i)
	}
	wg.Wait()
}
