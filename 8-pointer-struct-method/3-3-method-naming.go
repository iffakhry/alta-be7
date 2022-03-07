package main

import (
	"fmt"
	"math"
)

type Rect struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

//pakai pendekatan method
func (r Rect) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

//jika pakai pendekatan function
func LuasPersegi(panjang float64, lebar float64) float64 {
	return panjang * lebar
}

func LuasLingkaran(radius float64) float64 {
	return math.Pi * radius * radius
}

func main() {
	rect := Rect{5.0, 4.0}
	cir := Circle{5.0}
	fmt.Printf("Area of rectangle rect = %0.2f\n", rect.Area())
	fmt.Printf("Area of circle cir = %0.2f\n", cir.Area())
	fmt.Printf("Area of circle cir = %0.2f\n", LuasLingkaran(cir.radius))
	fmt.Printf("Area of circle cir = %0.2f\n", LuasPersegi(rect.width, rect.height))
}
