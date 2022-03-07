package main

import "fmt"

type Employee struct {
	name   string
	salary int
}

func (e *Employee) changeName(newName string) {
	(*e).name = newName
}

func (e *Employee) changeSalary(newSalary int) {
	(*e).salary = newSalary
}

func main() {
	e := Employee{
		name:   "Ross Geller",
		salary: 1200,
	}

	// e before name change
	fmt.Println("e before name change =", e)
	// create pointer to `e`
	ep := &e
	// change name
	ep.changeName("Monica Geller")
	// e after name change
	fmt.Println("e after name change =", e)
	ep.changeSalary(20000)
	// e after name change
	fmt.Println("e after name change =", e)
}
