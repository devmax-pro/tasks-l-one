// 6. Реализовать все возможные способы остановки выполнения горутины.
// 3. Использование sync.WaitGroup для ожидания завершения горутин

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик группы при выходе из функции.

	fmt.Printf("Worker %d starting\n", id)

	// Выполняем некоторую работу:
	time.Sleep(time.Second)

	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	// Запускаем несколько горутин и добавляем их в WaitGroup:
	for i := 1; i <= 5; i++ {
		wg.Add(1) // Увеличиваем счетчик группы.
		go worker(i, &wg)
	}

	// Ожидаем завершения всех горутин в WaitGroup:
	wg.Wait()

	fmt.Println("All workers completed")
}

// В этом примере создается 5 горутин, каждая из которых "работает" в течение одной секунды.
// WaitGroup используется для ожидания завершения всех горутин, прежде чем программа выведет "All workers completed".

// Обратите внимание, что WaitGroup сам по себе не останавливает выполнение горутин.
// Он просто предоставляет механизм для ожидания их завершения.
