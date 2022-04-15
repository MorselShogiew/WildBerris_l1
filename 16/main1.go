package main

import "fmt"

var slice2 = []int{5, 4, 1, 2, 0, -1, 1234, 32, 12, 2345, 99}

func quicksort(slice []int, left int, right int) {
	//Создаем копии пришедших переменных, с которыми будем манипулировать в дальнейшем.
	l := left
	r := right
	m := (left + right) / 2

	//Вычисляем 'центр', на который будем опираться. Берем значение ~центральной ячейки массива.
	center := slice[m]

	//Цикл, начинающий саму сортировку
	for l <= r {

		//Ищем значения больше 'центра'
		for slice[r] > center {
			r--
		}

		//Ищем значения меньше 'центра'
		for slice[l] < center {
			l++
		}

		//После прохода циклов проверяем счетчики циклов
		if l <= r {
			//И если условие true, то меняем ячейки друг с другом.
			//fmt.Printf("slice = %v; center = %v\n", slice, center)
			slice[r], slice[l] = slice[l], slice[r]
			l++
			r--
		}
	}

	if r > left {
		quicksort(slice, left, r)
	}

	if l < right {
		quicksort(slice, l, right)
	}
}

func main1() {
	fmt.Println("Started slice = ", slice2)
	quicksort(slice2, 0, len(slice2)-1)
	fmt.Println("Started slice = ", slice2)
}
