package main

import (
	"fmt"
	"time"
)

func f3(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}
func main() {

	f3("direct")
	go f3("goroutines")

	go func(arg string) {
		fmt.Println(arg)
	}("going")

	time.Sleep(time.Second)

	fmt.Println("done")
}
