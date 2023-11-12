// 25. Реализовать собственную функцию sleep
package main

import (
	"fmt"
	"time"
)

// CustomSleep пауза на заданное количество секунд
func CustomSleep(seconds int) {
	<-time.After(time.Duration(seconds) * time.Second)
}

func main() {
	fmt.Println("Sleeping for 3 seconds...")
	CustomSleep(3)
	fmt.Println("Woke up!")
}
