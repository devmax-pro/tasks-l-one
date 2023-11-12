// 14. Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel
// из переменной типа interface{}.

package main

import (
	"fmt"
)

func checkType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("The type is int, value: %d\n", v)
	case string:
		fmt.Printf("The type is string, value: %s\n", v)
	case bool:
		fmt.Printf("The type is bool, value: %t\n", v)
	case chan int:
		fmt.Printf("The type is channel int, value: %v\n", v)
	default:
		fmt.Printf("Unknown type\n")
	}
}

func main() {
	checkType(5)
	checkType("hello")
	checkType(true)
	checkType(make(chan int))
}
