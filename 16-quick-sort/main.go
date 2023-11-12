// 16. Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

package main

import (
	"fmt"
)

func quicksort(arr []int) []int {
	// Если слайс содержит ноль или один элемент, он уже отсортирован, возвращаем как есть.
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1

	// Выбираем опорный элемент
	pivotIndex := len(arr) / 2

	// Перемещаем опорный элемент в конец
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	// Размещаем элементы меньше опорного слева от него
	for i := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	// Ставим опорный элемент на его место, так, что слева от него теперь все элементы меньше его
	arr[left], arr[right] = arr[right], arr[left]

	// Применяем сортировку к левой и правой части массива
	quicksort(arr[:left])
	quicksort(arr[left+1:])

	return arr
}

func main() {
	arr := []int{3, 6, 8, 10, 1, 2, 1}
	fmt.Println("Original :", arr) // [3 6 8 10 1 2 1]

	sortedArr := quicksort(arr)
	fmt.Println("Sorted :", sortedArr) // [1 1 2 3 6 8 10]
}
