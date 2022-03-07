package main

import "fmt"

func fiboDPExtra(n int) int {
	if n <= 1 {
		return n
	}
	fiboI, fiboIMinSatu, fiboIMinDua := -1, 1, 0
	for i := 2; i <= n; i++ {
		fiboI = fiboIMinSatu + fiboIMinDua
		fiboIMinDua = fiboIMinSatu
		fiboIMinSatu = fiboI
	}
	return fiboI
}

func main() {
	fmt.Println(fiboDPExtra(6))
}
