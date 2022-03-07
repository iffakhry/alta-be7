package main

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
}

//function
func fullName(firstName string, lastName string) (fullName string) {
	fullName = firstName + " " + lastName
	return
}

func main() {
	e := Employee{
		FirstName: "Ross",
		LastName:  "Geller",
	}

	fmt.Println(fullName(e.FirstName, e.LastName))
}
