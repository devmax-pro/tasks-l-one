// 17. Реализовать бинарный поиск встроенными методами языка.
package main

import "fmt"

// BinarySearch ищет target в отсортированном слайсе nums и возвращает индекс, по которому он находится, или -1, если элемент не найден.
func BinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid // Целевой элемент найден.
		} else if nums[mid] < target {
			left = mid + 1 // Ищем в правой половине.
		} else {
			right = mid - 1 // Ищем в левой половине.
		}
	}

	return -1 // Элемент не найден.
}

func main() {
	nums := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20} // Отсортированный массив.
	target := 16

	index := BinarySearch(nums, target)
	if index != -1 {
		fmt.Printf("Элемент %d найден на индексе %d.\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден в массиве.\n", target)
	}
}
