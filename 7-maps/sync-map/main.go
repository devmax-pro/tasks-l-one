// 7. Реализовать конкурентную запись данных в map.
// 2. Используем sync.Map
package main

import (
	"fmt"
	"sync"
)

func main() {
	var sm sync.Map

	// Запускаем несколько горутин для записи в sync.Map
	var wg sync.WaitGroup

	// Количество горутин, которые будут записывать данные
	numGoroutines := 100
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(n int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", n)
			value := n
			sm.Store(key, value) // Безопасная запись в sync.Map
		}(i)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	sm.Range(func(key, value interface{}) bool {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
		return true
	})
}
