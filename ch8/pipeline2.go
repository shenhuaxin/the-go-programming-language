package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; ; x++ {
			time.Sleep(1 * time.Second)
			naturals <- x
		}
	}()

	// Squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
	}()

	// Printer (in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}

}