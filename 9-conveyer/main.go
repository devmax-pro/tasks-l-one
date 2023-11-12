// 9. Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
package main

import (
	"fmt"
)

func writeNumbers(numbers []int, in chan<- int) {
	for _, num := range numbers {
		in <- num
	}
	close(in)
}

func doubleNumbers(in <-chan int, out chan<- int) {
	for num := range in {
		out <- num * 2
	}
	close(out)
}

func printNumbers(out <-chan int) {
	for num := range out {
		fmt.Println(num)
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	in := make(chan int)  // канал для чисел
	out := make(chan int) // канал для результатов

	go writeNumbers(numbers, in)
	go doubleNumbers(in, out)

	printNumbers(out)
}
