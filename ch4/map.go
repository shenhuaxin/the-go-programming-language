package main

func main() {

	demo1()
	demo2()

	demo3()
}

func demo1() {
	ages := make(map[string]int)
	println(ages)
}

func demo2() {
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	println(ages)
}

func demo3() {
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	age, ok := ages["alice"]
	println(age, ok)

	age = ages["11"]
	println(age)
}
