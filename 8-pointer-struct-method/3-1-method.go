package main

import "fmt"

type Employee struct {
	FirstName, LastName string
}

//method
func (e Employee) fullName() string {
	return e.FirstName + " " + e.LastName
}

func (e Employee) sapaName() string {
	return "Hello! " + e.FirstName
}

func main() {
	emp := Employee{
		FirstName: "Ross",
		LastName:  "Geller",
	}
	fmt.Println(emp.fullName())
	fmt.Println(emp.sapaName())
}
