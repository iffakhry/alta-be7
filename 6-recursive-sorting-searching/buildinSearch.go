package main

import (
	"fmt"
	"sort"
)

func main() {
	elements := []int{12, 15, 15, 19, 24, 31, 53, 59, 60}
	value := 32
	findIndex := sort.SearchInts(elements, value)
	if value == elements[findIndex] {
		fmt.Println("value found in elemenets", findIndex)
	} else {
		fmt.Println("value not found in elemenets", findIndex)
	}
}
