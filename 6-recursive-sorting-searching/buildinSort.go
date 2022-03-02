package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// An example of sorting `int`s.
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	// We can also use `sort` to check if a slice is
	// already in sorted order.
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)

	data := []int{2, 5, 7, 3, 1, 9, 4}
	sort.SliceStable(data, func(i, j int) bool {
		return data[i] > data[j]
	})
	fmt.Println("data sorted:", data)

	dataString := []string{"c", "a", "b"}
	sort.SliceStable(dataString, func(i, j int) bool {
		return dataString[i] > dataString[j]
	})
	fmt.Println("dataString sorted:", dataString)
	hasil := strings.Join(dataString, "")
	fmt.Println(hasil)

}
