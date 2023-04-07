package main

func vals() (success bool, res string) {

	return false, "错误原因"
}
func main() {

	success, res := vals()
	println(success, res)

}
