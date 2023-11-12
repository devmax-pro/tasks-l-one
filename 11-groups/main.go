// 11. Реализовать пересечение двух неупорядоченных множеств.
package main

import (
	"fmt"
)

func intersection(set1, set2 []int) []int {
	m := make(map[int]bool)
	var intersectionSet []int

	for _, elem := range set1 {
		m[elem] = true
	}

	for _, elem := range set2 {
		if _, exists := m[elem]; exists {
			intersectionSet = append(intersectionSet, elem) // Если элемент присутствует в двух множествах, добавляем его
			delete(m, elem)                                 // Чтобы избежать дублирования в результате, удаляем элемент
		}
	}

	return intersectionSet
}

func main() {

	// Пример использования:
	setA := []int{1, 3, 5, 7, 9}
	setB := []int{0, 2, 3, 4, 5, 6, 7}

	result := intersection(setA, setB)
	fmt.Println("Пересечение: ", result) // Вывод: [3 5 7]

}
