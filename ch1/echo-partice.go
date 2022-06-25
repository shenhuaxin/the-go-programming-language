package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(os.Args[0])

	for i := 1; i < len(os.Args); i++ {
		println(strconv.Itoa(i) + " " + os.Args[i])
	}

}
