// 21. Реализовать паттерн «адаптер» на любом примере.
package main

import "fmt"

// Printer интерфейс, который мы будем реализовывать
type Printer interface {
	Print()
}

// LegacyPrinter структура, которую будем адаптировать
type LegacyPrinter struct{}

// AdaptedPrint - метод адаптируемой структуры
func (p *LegacyPrinter) AdaptedPrint() {
	fmt.Println("Message from Legacy Printer")
}

// ConcretePrinter - адаптер для LegacyPrinter
type ConcretePrinter struct {
	oldPrinter *LegacyPrinter
}

func NewConcretePrinter(oldPrinter *LegacyPrinter) Printer {
	return &ConcretePrinter{
		oldPrinter: oldPrinter,
	}
}

// Print - метод адаптера
func (p *ConcretePrinter) Print() {
	fmt.Println("Message from Concrete Printer")
	p.oldPrinter.AdaptedPrint()
}

func main() {
	myNewPrinter := ConcretePrinter{&LegacyPrinter{}}
	myNewPrinter.Print()
}
