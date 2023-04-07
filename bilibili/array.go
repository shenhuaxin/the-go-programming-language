package main

import "fmt"

func main() {

	array1 := [...]string{
		"1", "2", "3", "4",
	}

	array2 := array1

	array1[1] = "22"

	fmt.Println(array1)
	fmt.Println(array2)
}
