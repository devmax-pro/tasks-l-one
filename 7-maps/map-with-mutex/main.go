// 7. Реализовать конкурентную запись данных в map.

// 7.1 Используем Mutex
package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu    sync.Mutex
	store map[string]interface{}
}

// Set устанавливает значение по ключу с блокировкой для синхронизации
func (sm *SafeMap) Set(key string, value interface{}) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.store[key] = value
}

// Get возвращает значение по ключу с блокировкой
func (sm *SafeMap) Get(key string) (interface{}, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	val, ok := sm.store[key]
	return val, ok
}

func main() {
	sm := SafeMap{
		store: make(map[string]interface{}),
	}

	// Количество горутин
	numGoroutines := 100

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(n int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", n)
			sm.Set(key, n)
		}(i)
	}

	// Ожидание завершения всех горутин
	wg.Wait()

	for i := 0; i < numGoroutines; i++ {
		key := fmt.Sprintf("key%d", i)
		if value, ok := sm.Get(key); ok {
			fmt.Printf("Key: %s, Value: %d\n", key, value)
		}
	}
}
