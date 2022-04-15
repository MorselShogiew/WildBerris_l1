package main

import (
	"fmt"
	"time"
)

func Sleep(num int) {
	<-time.After(time.Second * time.Duration(num))
}
func main() {
	fmt.Println("asrqwrqf")
	Sleep(3)
	fmt.Println("asrqwrqaegewgweg")
}
