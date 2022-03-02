package main

import "fmt"

func main() {

	kata := "hello"
	fmt.Println("panjang kata:", len(kata))

	// for i := 0; i <= 5; i++ {
	// 	fmt.Println(i)
	// 	fmt.Println("halo dunia")
	// }

	sum := 1
	var result int
	for sum < 10 {
		sum += sum
		result = sum + sum
		fmt.Println(sum)
		// sum = sum + sum
		/*
			sum += 1
			sum = sum + 1

				1 .. sum = 1
				2 .. sum = 1 + 1 = 2
				3 .. sum = 2 + 2 = 4
				4 .. sum = 4 + 4 = 8
				5 .. sum = 8 + 8 = 16
		*/
	}
	fmt.Println(sum)
	fmt.Println(result)
}
