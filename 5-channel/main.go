// 5. Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться.

package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int)

	// Отправляем значения в канал
	go func() {
		for i := 0; ; i++ {
			channel <- i
			time.Sleep(time.Second)
		}
	}()

	// Читаем значения из канала
	go func() {
		for {
			fmt.Println(<-channel)
		}
	}()

	// Ждем 5 секунд перед завершением программы
	time.Sleep(5 * time.Second)
}
