package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string
	Age  int
}

type response1 struct {
	Page   int
	Fruits []string
}

//func newPerson(name string) *person {
//	p := person{name: name}
//	p.age = 23
//	return &p
//}
func main() {
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	p := &person{"shenhuaxin", 13}

	marshal, _ := json.Marshal(p)

	fmt.Println(string(marshal))
}
