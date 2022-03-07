package main

import (
	"fmt"
	"math"
)

type calculate interface {
	large() int
	keliling() int
	// sepertiga() int
}

type square struct {
	side int
}

type circle struct {
	radius float64
}

func (s square) large() int {
	return s.side * s.side
}

func (s square) keliling() int {
	return 4 * s.side
}

// func (s square) sepertiga() int {
// 	return 1
// }

func (c circle) large() int {
	result := math.Pi * c.radius * c.radius
	return int(result)
}

func (c circle) keliling() int {
	result := 2 * math.Pi * c.radius
	return int(result)
}

func main() {
	var dimResult calculate
	dimResult = square{10}
	fmt.Println("large square :", dimResult.large())
	fmt.Println("keliling square :", dimResult.keliling())

	var dimCircleResult calculate
	dimCircleResult = circle{10}
	fmt.Println("large circle :", dimCircleResult.large())
	fmt.Println("keliling circle :", dimCircleResult.keliling())
}
