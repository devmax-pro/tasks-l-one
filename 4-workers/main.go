// 4. Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают
// произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, id int, data <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done(): // если контекст был отменен, завершаем горутину
			fmt.Printf("Worker %d: stopping\n", id)
			return
		case value := <-data: // читаем данные из канала
			fmt.Printf("Worker %d: received %d\n", id, value)
		}
	}
}

func main() {
	var workersNum int

	fmt.Println("Enter the number of workers:")
	_, err := fmt.Scan(&workersNum)
	if err != nil {
		fmt.Printf("Error reading number of workers: %v\n", err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	data := make(chan int)
	wg := &sync.WaitGroup{}

	// Создаем и запускаем воркеры
	for i := 0; i < workersNum; i++ {
		wg.Add(1)
		go worker(ctx, i+1, data, wg)
	}

	// Обработка сигналов для корректного завершения работы при нажатии Ctrl+C
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signals // ожидаем сигнал
		fmt.Println("\nReceived an interrupt, stopping workers...")
		cancel()    // отменяем контекст
		close(data) // закрываем канал
	}()

	// Эмуляция постоянной записи данных в канал
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				data <- rand.Intn(1000) // записываем случайное значение в канал
				time.Sleep(time.Second) // пауза для наглядности
			}
		}
	}()

	// Ожидаем завершения всех воркеров
	wg.Wait()
	fmt.Println("All workers stopped.")
}
