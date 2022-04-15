package main

import (
	"fmt"
	"sync"
	"time"
)

func MapIm(ma map[int]int, m *sync.Mutex, value int) {
	for key := 0; key < 5; key++ {
		m.Lock()
		ma[key] = value
		m.Unlock()
		fmt.Printf("map [%d]=%d\n", key, value)
	}
}
func main() {
	map1 := make(map[int]int)
	var m sync.Mutex
	for i := 0; i < 3; i++ {
		go MapIm(map1, &m, i)
	}
	time.Sleep(time.Second * 1)
	for key, value := range map1 {
		fmt.Printf("key: %v\tvalue: %v\n", key, value)
	}
}
