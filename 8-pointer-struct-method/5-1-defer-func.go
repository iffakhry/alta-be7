package main

import "fmt"

/*
Stack = LIFO
queue = FIFO
*/
func main() {
	fmt.Println("awalan")
	defer func() {
		fmt.Println("Later")
	}()
	defer func() {
		fmt.Println("kemudian")
	}()
	defer func() {
		fmt.Println("selanjutnya")
	}()

	fmt.Println("First")
	fmt.Println("second")

}

// func sum(numbers ...int) int {

// 	var total int = 0
// 	for _, number := range numbers {
// 		total += number
// 	}
// 	return total
// }

// func main() {
// 	avg := sum(2, 4, 3, 5, 7, 8, 1, 2)
// 	fmt.Println(avg)
// }
