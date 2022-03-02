package main

import "fmt"

func BinarySearch(data []int, x int) int {
	// var hasil int
	var kiri = 0
	var kanan = len(data) - 1
	var counter int
	for kiri <= kanan {
		counter++
		var mid = (kiri + kanan) / 2
		fmt.Println("mid", mid)
		if data[mid] == x {
			return mid
		} else if x < data[mid] {
			kanan = mid - 1
		} else if x > data[mid] {
			kiri = mid + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(BinarySearch([]int{1, 1, 3, 5, 5, 6, 7, 8}, 7))
}
