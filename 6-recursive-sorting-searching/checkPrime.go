package main

import (
	"fmt"
	"math"
)

func checkPrime(n int) bool {
	if n < 2 {
		return false
	}
	sqrtNumber := int(math.Sqrt(float64(n)))

	for i := 2; i <= sqrtNumber; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(checkPrime(1000000007))
	// fmt.Println(checkPrime(6))
}
