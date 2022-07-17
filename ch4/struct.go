package main

import "time"

type Employee struct {
	ID        int
	Name      string
	Address   string
	Dob       time.Time
	Position  string
	Salary    int
	ManagerId int
}

func main() {
	var dilbert Employee
	dilbert.Salary -= 5000

	println(dilbert.Salary)

	position := &dilbert.Position
	*position = "Senior " + *position
	println(dilbert.Position)
}
