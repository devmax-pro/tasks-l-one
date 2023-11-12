// 1. Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от
// родительской структуры Human (аналог наследования).
package main

import (
	"fmt"
)

// Human - базовая структура с полями и методами
type Human struct {
	Name string
	Age  int
}

// Speak - метод, который позволяет человеку говорить
func (h *Human) Speak() {
	fmt.Printf("My name is %s and I am %d years old.\n", h.Name, h.Age)
}

// Action - производная структура, которая встраивает Human
type Action struct {
	Human // Встраивание структуры Human
}

func (a *Action) Jump() {
	fmt.Printf("%s jumps!\n", a.Name) // Обращение к полю Name структуры Human
}

func main() {
	action := Action{
		Human: Human{
			Name: "John",
			Age:  30,
		},
	}

	// Вызов метода Speak для структуры Human
	action.Speak()

	// Вызов метода Jump для структуры Action
	action.Jump()
}
