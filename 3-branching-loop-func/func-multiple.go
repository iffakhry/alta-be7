package main

import (
	"fmt"
	"math"
)

// singe return value
func calculateSquare(side int) int {
	return side * side
}

// multiple return value
func calculateCircle(diameter float64) (float64, float64) {
	var keliling = math.Pi * math.Pow(diameter/2, 2)
	var luas = math.Pi * diameter
	// return 2 value
	return keliling, luas
}

func calculateTriangle(a string, b int, c int) int {
	return 1
}

func main() {
	calculateTriangle("hallo", 3, 7)
	var side = 5
	wide := calculateSquare(side)
	fmt.Printf("Luas persegi empat: %d \n\n", wide)

	var side2 = 10
	wide2 := calculateSquare(side2)
	fmt.Printf("Luas persegi empat ke 2: %d \n\n", wide2)

	var diameter float64 = 15
	var keliling1, luas1 float64
	keliling1, _ = calculateCircle(diameter)
	fmt.Printf("Luas lingkaran: %.2f \n", luas1)
	fmt.Printf("Keliling lingkaran: %.2f \n", keliling1)
}
