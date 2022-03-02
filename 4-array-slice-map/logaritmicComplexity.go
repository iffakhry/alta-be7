package main

import "fmt"

func logarithmic(n int) int {
	var result int = 0
	for n > 1 {
		n /= 2      // n = n/2
		result += 1 // result = result + 1
		fmt.Println("result:", result, " n: ", n)
	}
	return result
}

/*
2^x = 32
2^5 = 32
*/
func main() {
	result := logarithmic(100)
	fmt.Println(result)
}
