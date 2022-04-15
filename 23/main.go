package main

import "fmt"

var (
	slice = []int{0, 1, 2, 3, 4, 5, 6, 7}
	index = 3
)

func main() {
	fmt.Println("Started slice = ", slice)
	//сдвиг a[i+1:] влево на один индекс
	copy(slice[index:], slice[index+1:])

	fmt.Println("Finished slice = ", slice)

	//меняет порядок
	//main1()
	main2()
}
