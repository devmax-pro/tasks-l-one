// 12. Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

// В этом примере struct{} используется в качестве значения в Map для имитации множеств.
// В принципе, можно использовать bool или int, но пустая структура не занимает дополнительной памяти.
package main

import "fmt"

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]struct{})

	for _, element := range sequence {
		set[element] = struct{}{}
	}

	for element := range set {
		fmt.Println(element)
	}
}
