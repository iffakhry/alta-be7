package main

import "fmt"

var count int

func fibonacci(n int) int {
	count++
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// 0 1 1 2 3 5 8 13 ....
func main() {
	var bil = 6
	fmt.Println(fibonacci(bil)) //8
	// fmt.Println(fibonacci(50))
	fmt.Println(count)
}
