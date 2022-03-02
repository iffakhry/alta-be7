package main

import "fmt"

func main() {
	// var temp = []int{}

	var colors = []string{"red", "green", "yellow"}
	fmt.Println("sebelum", len(colors))
	colors = append(colors, "purple")
	colors = append(colors, "white")
	for _, v := range colors {
		fmt.Println(v)
	}
	fmt.Println(len(colors))
	fmt.Println(colors[2])

	var primes []int
	fmt.Printf("s = %v, len = %d, cap = %d\n", primes, len(primes), cap(primes))

	if primes == nil {
		fmt.Println("s is nil")
	}

	//menggabungkan 2 slice
	slice1 := []int{1, 2, 3}
	slice2 := []int{3, 4, 5, 6}
	slice1 = append(slice1, slice2...)
	fmt.Println("slice1", slice1)
}
