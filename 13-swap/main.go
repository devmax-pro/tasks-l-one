// 13. Поменять местами два числа без создания временной переменной
package main

import "fmt"

func main() {
	a := 5
	b := 10

	fmt.Printf("Before : a = %d, b = %d\n", a, b)

	a, b = b, a

	fmt.Printf("After : a = %d, b = %d\n", a, b)
}
