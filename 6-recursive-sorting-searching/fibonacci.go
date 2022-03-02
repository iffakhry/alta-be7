package main

import "fmt"

// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144,
func fibonacci(n int) int {
	//base case
	if n <= 1 {
		return n
	}
	//recursive case
	return fibonacci(n-1) + fibonacci(n-2)
	//      fibo(3-1) + fibo(3-2) --> fibo(2)+ fibo(1)
	// 		fibo(1)+fibo(0) + fibo(1)
	// 		1+0+1
}

/*
fibo(3) + fibo(2)
fibo(2) + fibo(1) + fibo(1)+fibo(0)
fibo(1)+fibo(0) + 1 + 1 + 0
1 + 0 + 1 + 1 + 0 = 3
*/

func main() {
	fmt.Println(fibonacci(2)) //1
	fmt.Println(fibonacci(4)) //3
	fmt.Println(fibonacci(9)) //34
}
