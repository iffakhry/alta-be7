package main

import "fmt"

var memoTopDown = map[int]int{}
var count int

func fiboTopDown(n int) int {
	count++
	//proses membaca memo -> jika sudah pernah dihitung sebelumnya, maka langsung baca hasilnya dari kamus
	value, sudahDihitung := memoTopDown[n]
	if sudahDihitung {
		return value
	}

	//jika belum pernah dihitung, maka lakukan perhitungan, dan simpan hasilnya kedalam kamus.
	if n <= 1 {
		memoTopDown[n] = n
	} else {
		memoTopDown[n] = fiboTopDown(n-1) + fiboTopDown(n-2)
	}
	fmt.Println("count", count)
	return memoTopDown[n]
}

func main() {
	fmt.Println(fiboTopDown(6))
	// fmt.Println(fiboTopDown(30))
	// fmt.Println(count)
	// fmt.Println(fiboTopDown(50))
}
