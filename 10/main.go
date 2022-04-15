package main

import "fmt"

func main() {
	weath := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	group_k10 := make(map[int][]float32)
	var d int
	for _, num := range weath {
		if num < 0 {
			d = int((num-10)/10) * 10
			group_k10[d] = append(group_k10[d], num)
		} else {
			d = int(num/10) * 10
			group_k10[d] = append(group_k10[d], num)
		}

	}

	fmt.Println(group_k10)
}
