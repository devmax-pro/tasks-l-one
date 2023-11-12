// 3. Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов с использованием конкурентных вычислений.
package main

import (
	"fmt"
	"sync"
)

// Вычисляем квадрат числа и отправляем результат в канал
func square(wg *sync.WaitGroup, num int, ch chan int) {
	defer wg.Done() // Уменьшаем счетчик при завершении горутины
	ch <- num * num
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	// Создаем канал для передачи результатов
	ch := make(chan int, len(numbers))

	// Создаем WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Запускаем горутину для каждого числа в массиве
	for _, number := range numbers {
		wg.Add(1) // Увеличиваем счетчик группы ожидания
		go square(&wg, number, ch)
	}

	// Ждем завершения всех горутин
	wg.Wait()

	// Закрываем канал, так как все горутины завершили свою работу и больше ничего в него писать не будут
	close(ch)

	// Суммируем результаты из канала
	sum := 0
	for sqr := range ch {
		sum += sqr
	}

	fmt.Printf("Sum of squares: %d\n", sum)
}
