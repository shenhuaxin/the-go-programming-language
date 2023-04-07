package main

import "fmt"

func sum(num ...int) int {
	total := 0
	for i := range num {
		total += i
	}
	return total
}
func main() {

	fmt.Println(sum(1, 2, 3, 4, 5, 6))
}
