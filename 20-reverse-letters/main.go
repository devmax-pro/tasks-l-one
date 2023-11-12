// 20. Разработать программу, которая переворачивает слова в строке. Пример: «snow dog sun — sun dog snow».
package main

import (
	"fmt"
	"strings"
)

func ReverseWords(s string) string {
	words := strings.Fields(s) // Возвращает подстроки, которые являются «словами» в исходной строке
	for i := 0; i < len(words)/2; i++ {
		j := len(words) - i - 1
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	s := "snow dog sun"
	fmt.Println(ReverseWords(s)) // "sun dog snow"
}
