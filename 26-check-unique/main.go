// 26. Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.
// Например:
// abcd — true abCdefAaf — false aabcd — false

package main

import (
	"fmt"
	"strings"
)

func IsUnique(str string) bool {
	str = strings.ToLower(str)
	chars := make(map[rune]bool)

	// Итерация по символам строки
	for _, char := range str {
		if _, exists := chars[char]; exists {
			// Если символ уже есть в мапе, возвращаем false
			return false
		}
		chars[char] = true
	}
	// Все символы уникальны, возвращаем true
	return true
}

func main() {
	fmt.Println(IsUnique("abcd"))      // true
	fmt.Println(IsUnique("abCdefAaf")) // false
	fmt.Println(IsUnique("aabcd"))     // false
}
