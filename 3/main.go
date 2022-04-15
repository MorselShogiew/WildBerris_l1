package main

import (
	"fmt"
	"sync"
)

var (
	arr1 = []int{2, 4, 6, 8, 10}
	ans  int
)

func Square1(v int, wg *sync.WaitGroup, m *sync.Mutex) {
	sqv := v * v
	m.Lock()
	defer m.Unlock()
	ans += sqv
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex
	for _, val := range arr1 {
		wg.Add(1)
		go Square1(val, &wg, &m)
	}
	wg.Wait()
	fmt.Printf("ans = %v\n", ans)

	//main1()
}
