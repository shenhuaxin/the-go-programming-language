package main

func main() {

	println(gcd(7, 9))

	println(fib(10))

	var m = make(map[int]int)
	m[1] = 3
	v, ok := m[1]
	println(v)
	println(ok)

	b := 1 == 1
	println(b)
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
