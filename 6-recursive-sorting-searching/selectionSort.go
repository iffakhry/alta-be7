package main

import "fmt"

func selectionSort(elements []int) []int {
	n := len(elements)
	for k := 0; k < n; k++ {
		minimal := k
		for j := k + 1; j < n; j++ {
			if elements[j] < elements[minimal] {
				minimal = j
			}
		}
		//tukar posisi data
		elements[k], elements[minimal] = elements[minimal], elements[k]
	}
	return elements
}

func main() {
	/*
		var a2, b2 int
		a2 = 1
		b2 = 2
		//tukar data
		a2, b2 = b2, a2
		fmt.Println("a2:", a2, "b2:", b2)
	*/
	fmt.Println(selectionSort([]int{4, 7, 2, 9, 1, 5}))
}
