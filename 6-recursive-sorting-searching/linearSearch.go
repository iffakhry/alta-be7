package main

import "fmt"

func linierSearch(elements []int, x int) int {
	for i := 0; i < len(elements); i++ {
		if elements[i] == x {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(linierSearch([]int{2, 4, 5, 7, 9, 3}, 10))
}
