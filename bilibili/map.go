package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	ages := map[string]int{}

	ages["shx"] = 25
	ages["xxx"] = 26

	b, err := json.Marshal(ages)
	if err != nil {
		return
	}
	fmt.Println(string(b))

	for k, v := range ages {
		println(k, v)
	}

	delete(ages, "xxx")

	k, v := ages["shx"]

	println(k, v)

	fmt.Println("ages", ages)
}
