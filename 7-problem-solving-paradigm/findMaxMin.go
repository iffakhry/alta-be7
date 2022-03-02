package main

import "fmt"

func MinMax(data []int) (int, int) {
	var max, min int
	max = data[0]
	min = data[0]
	for i := 1; i < len(data); i++ {
		if data[i] < min {
			min = data[i]
		}
		if data[i] > max {
			max = data[i]
		}
	}
	return min, max
}

func main() {
	data := []int{1, 8, 5, 9, 2, 4, 3}
	min, max := MinMax(data)
	fmt.Println(min, max)
}
