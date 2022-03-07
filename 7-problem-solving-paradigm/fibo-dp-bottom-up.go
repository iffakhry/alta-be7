package main

import "fmt"

func fiboBottomUp(n int) int {
	var count int
	var memoBottomUp map[int]int = map[int]int{}
	for i := 0; i <= n; i++ {
		if i <= 1 {
			memoBottomUp[i] = i
		} else {
			memoBottomUp[i] = memoBottomUp[i-1] + memoBottomUp[i-2]
		}
		count++
	}
	fmt.Println("count:", count)
	return memoBottomUp[n]
}

func main() {
	fmt.Println(fiboBottomUp(6)) //8
	// fmt.Println(fiboBottomUp(50)) //12586269025
}
