/*
Задача:
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

	var justString string
	func someFunc() {
	   v := createHugeString(1 << 10)
	   justString = v[:100]
	}
	func main() {
	   someFunc()
	}

Проблемы:

1. Утечка памяти:
	Строка v, созданная функцией createHugeString(1 << 10), представляет собой большую строку (1024 символа).
Подстрока justString = v[:100] ссылается на первые 100 символов этой строки. В Go строки неизменяемы,
и подстрока фактически ссылается на исходные байты строки v, что означает, что пока justString находится в области видимости
и не уничтожается сборщиком мусора, вся огромная строка v также остается в памяти, что приводит к утечке.
Решение: Чтобы избежать утечки памяти, можно скопировать нужную подстроку в новую строку, которая не будет ссылаться на исходные данные большой строки.

Начиная с Go 1.18, стандартная библиотека также включает решение с strings.Clone(s string), которое возвращает новую копию строки: `result := strings.Clone(s[:100])`
Вызов strings.Clone создает копию s[:100] в новом месте памяти, предотвращая утечку памяти.

2. Проблема при использовании не ASCII символов:
	В Go строка представляет собой последовательность байтов, и каждый символ ASCII кодируется одним байтом.
Однако Unicode символы, находящиеся за пределами ASCII, например символы кириллицы, иероглифы и др. кодируются несколькими байтами в UTF-8.
При выполнении операции среза v[:100], как в изначальном примере, мы рискуем разделить один из этих многобайтовых символов пополам,
что приведет к созданию недопустимой UTF-8 последовательности.
Решение: Чтобы корректно работать с не ASCII символами, нужно конвертировать строку в слайс рун и затем сделать операцию среза

3. Глобальная переменная:
	Переменная justString может быть потенциальным источником проблем в многопоточной среде и с точки зрения дизайна кода.
Глобальные переменные делают состояние программы менее предсказуемым и могут привести к усложнению тестирования и отладки кода.
Решение: Лучше избегать глобальных переменных, если это возможно, и вместо этого возвращать значения из функций и передавать их туда, где они нужны.
*/

// Вот пример исправленной реализации:
package main

import (
	"fmt"
)

func createHugeString(size int) string {
	var str string
	// Просто пример, добавляющий символ 'a' в строку
	for i := 0; i < size; i++ {
		str += "a"
	}
	return str
}

func someFunc() string {
	v := createHugeString(1 << 10)
	// Конвертируем строку в слайс рун
	return string([]rune(v)[:100])
}

func main() {
	// Используем justString здесь
	justString := someFunc()
	fmt.Println(justString)
}
