// 2. Написать программу, которая конкурентно рассчитает значение квадратов чисел,
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 3, 4, 6, 8, 10}
	squares := make([]int, len(numbers))
	var wg sync.WaitGroup

	for i, number := range numbers {
		wg.Add(1) // Увеличиваем счетчик

		go func(i, number int) {
			defer wg.Done() // Уменьшаем счетчик
			squares[i] = number * number
		}(i, number)
	}

	wg.Wait() // Ожидаем завершения всех горутин

	for i, square := range squares {
		fmt.Printf("Square of number %d equals %d\n", numbers[i], square)
	}
}
