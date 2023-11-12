// 6. Реализовать все возможные способы остановки выполнения горутины.
// 2. Использование контекста
package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Got the signal to stop!")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go doWork(ctx)

	time.Sleep(2 * time.Second)        // Подождать некоторое время перед остановкой горутины
	cancel()                           // Отменить контекст
	time.Sleep(500 * time.Millisecond) // Дать время горутине на завершение
}
