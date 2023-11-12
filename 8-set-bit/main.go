// 8. Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

package main

import "fmt"

func setBit(n int64, i uint) int64 {
	return n | (1 << i)
}

func clearBit(n int64, i uint) int64 {
	return n &^ (1 << i)
}

func main() {
	var originalNum int64 = 8 // В двоичном виде 1000
	fmt.Printf("Исходное число %d: %b\n", originalNum, originalNum)

	// Установим 1-й бит в 1
	n := setBit(originalNum, 1) // 1010
	fmt.Printf("Установка 1-го бита в 1 для числа %d: %d - %b\n", originalNum, n, n)

	originalNum = 34 // В двоичном виде 100010
	fmt.Printf("Исходное число %d: %b\n", originalNum, originalNum)

	// Установим 1-й бит в 0
	n = clearBit(originalNum, 1) // 100000
	fmt.Printf("Сброс 1-го бита в 0 для числа %d: %d - %b\n", originalNum, n, n)
}
