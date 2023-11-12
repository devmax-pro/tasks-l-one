// 24. Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point
// с инкапсулированными параметрами x, y и конструктором.

package main

import (
	"fmt"
	"math"
)

// Point - структура с приватными полями x и y
type Point struct {
	x float64
	y float64
}

// NewPoint - конструктор для создания новой точки
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// Distance - вычисляет расстояние между двумя точками по формуле √((x2−x1)²+(y2−y1)²)
func (p *Point) Distance(other *Point) float64 {
	return math.Sqrt(math.Pow(p.x-other.x, 2) + math.Pow(p.y-other.y, 2))
}

func main() {
	// Создание двух точек
	p1 := NewPoint(1, 2)
	p2 := NewPoint(4, 6)

	// Вычисление расстояния между точками
	distance := p1.Distance(p2)

	// Вывод результата
	fmt.Printf("Расстояние между двумя точками: %v\n", distance)
}
