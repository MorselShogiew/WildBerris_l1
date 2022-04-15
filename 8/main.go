package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int64
	i = 9223372036854775800
	var val int64
	var pos int64
	fmt.Println(strconv.FormatInt(i, 2))
	fmt.Scan(&val)
	fmt.Scan(&pos)
	if val == 0 {
		fmt.Println(strconv.FormatInt(i&(9223372036854775807-1<<pos), 2))
	}
	if val == 1 {
		fmt.Println(strconv.FormatInt(i|1<<pos, 2))
	}
}
