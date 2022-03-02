package main

import "fmt"

func main() {

	N := 5
	// for i := 0; i < N; i++ {
	// 	for j := 0; j <= i; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	for i := N; i >= 0; i = i - 3 {
		fmt.Println(i)
	}

}
