package main

import "learn-go/ch2/popcount"

func main() {

	println(popcount.Popcount(7))

	println(popcount.PopcountLoop(7))

	println(popcount.PopcountShifting(7))

}
