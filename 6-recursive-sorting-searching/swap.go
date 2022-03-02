package main

import "fmt"

func main() {
	var a, b, c int
	a = 1
	b = 2

	//tukar data
	c = a
	a = b
	b = c

	fmt.Println("a:", a, "b:", b)

	var a2, b2 int
	a2 = 1
	b2 = 2
	//tukar data
	a2, b2 = b2, a2
	fmt.Println("a2:", a2, "b2:", b2)

}
