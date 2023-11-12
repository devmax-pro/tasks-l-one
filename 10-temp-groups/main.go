// 10. Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
package main

import (
	"fmt"
	"math"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int][]float64)

	for _, temp := range temperatures {

		var bin int
		if temp < 0 {
			// Для отрицательных значений округляем вверх до ближайшей десятки
			bin = int(math.Ceil(temp/10.0)) * 10
		} else {
			// Для положительных значений округляем вниз до ближайшей десятки
			bin = int(math.Floor(temp/10.0)) * 10
		}

		groups[bin] = append(groups[bin], temp)
	}
	for bin, temps := range groups {
		fmt.Printf("%d: %v\n", bin, temps)
	}
}
