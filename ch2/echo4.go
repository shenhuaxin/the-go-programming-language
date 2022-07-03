package main

import (
	"flag"
	"fmt"
	"os"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {

	//flag.Parse()
	//fmt.Print(strings.Join(flag.Args(), *sep))
	//
	//if !*n {
	//	fmt.Println()
	//}
	//
	//println("=======================>")
	//declare()
	//println("=======================>")
	//ponit()
	//

	newdemo()
}

func newdemo() {
	p := new(int)
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)
}

func ponit() {
	// 指针
	x := 1
	p := &x
	fmt.Println(*p)

	*p = 2
	fmt.Println(x)

	var x1, y int
	fmt.Println(x1)
	fmt.Println(&x1 == &x1, &x1 == &y, &x1 == nil)

	v := 1
	incr(&v)
	fmt.Println(incr(&v))
}

func incr(p *int) int {
	*p++
	return *p
}

func declare() {
	// 变量声明
	var s string = "1"
	print(s)
	// 简短变量声明
	f, err := os.Open("infile")
	f, err = os.Create("outfile") // compile error: no new variables

	println(f)
	println(err)
}
