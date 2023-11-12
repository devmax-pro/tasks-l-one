// 23. Удалить i-ый элемент из слайса.
package main

import "fmt"

func Remove(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(Remove(slice, 2)) // [1 2 4 5]
}
