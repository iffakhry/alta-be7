package main

import (
	"fmt"
	"sort"
)

func coinChange(money int, coins []int) []int {
	result := []int{}

	sort.SliceStable(coins, func(i, j int) bool {
		return coins[i] > coins[j]
	})
	// fmt.Println(coin)
	for _, coin := range coins {
		if money >= coin {
			for money >= coin {
				result = append(result, coin)
				money -= coin
			}
		} else {
			continue
		}
	}
	return result
}

func main() {
	/*
		42 = 25, 10, 5, 1,1
		42 = 1,1,1,1,1,1,1,1, .... 42x
		42 = 10,10,10,10,1,1
		42 = 25,5,5,5,1,1
	*/
	// coinValue = {10, 25, 5, 1}
	coinValue := []int{10, 25, 5, 1}
	fmt.Println(coinChange(42, coinValue))
	fmt.Println(coinChange(56, coinValue))
	fmt.Println(coinChange(16, coinValue))
}
