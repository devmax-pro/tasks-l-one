// 19. Разработать программу, которая переворачивает подаваемую на вход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

package main

import (
	"fmt"
)

func Reverse(s string) string {
	b := []rune(s)
	for i := 0; i < len(b)/2; i++ {
		j := len(b) - i - 1
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func main() {
	s := "главрыба😀"
	fmt.Println(Reverse(s)) // "😀абырвалг"
}
