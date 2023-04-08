package main

type rect struct {
	width, length int
}

func area(r *rect) int {
	return r.width * r.length
}

func (r *rect) area() int {
	return r.width * r.length
}

func (r rect) perim() int {
	return 2*r.width + 2*r.length
}

func main() {
	println(area(&rect{10, 20}))

	r := rect{10, 20}
	println(r.area())
	println(r.perim())
}
