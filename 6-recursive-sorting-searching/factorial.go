package main

import "fmt"

func factorialLoop(n int) int {
	var hasil int = 1
	for i := n; i > 0; i-- {
		hasil = hasil * i
	}
	return hasil
}

func factorialRecursive(n int) int {
	//base case
	if n == 1 {
		return 1
	}
	//recurrent relation
	return n * factorialRecursive(n-1)
}

/*
n=1 --> 1
n=5 --> 5 * factorialRecursive(4)
n=4 --> 5 * 4 * factorialRecursive(3)
n=3 --> 5 * 4 * 3 * factorialRecursive(2)
n=2 --> 5 * 4 * 3 * 2 * factorialRecursive(1)
n=1 --> 5 * 4 * 3 * 2 * 1
*/

func main() {
	fmt.Println(factorialLoop(5))      //120
	fmt.Println(factorialRecursive(5)) //120
}
