package main

import "fmt"

//n^2
func quadratic(n int) int {
	var result int = 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result += 1
		}
	}
	return result
}

//n^3
func quadraticn3(n int) int {
	var result int = 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				result += 1
			}
		}
	}
	return result
}

func main() {
	result := quadratic(10)
	fmt.Println(result)
}
