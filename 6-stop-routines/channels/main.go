// 6. Реализовать все возможные способы остановки выполнения горутины.
// 1. Использование каналов для сигнала завершения

package main

import (
	"fmt"
	"time"
)

func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("Got the signal to stop!")
			return
		default:
			// Делаем некоторую работу
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	done := make(chan bool)
	go doWork(done)

	time.Sleep(2 * time.Second)        // Подождать некоторое время перед остановкой горутины
	done <- true                       // Отправить сигнал для остановки горутины
	time.Sleep(500 * time.Millisecond) // Дать время горутине на завершение
}
