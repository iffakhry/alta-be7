package main

import "fmt"

//fungsi untuk mencari bilangan 0 dalam sebuah slice
func linier(n int, A []int) int {
	var counter int
	// if len(A) >= 4 {
	// 	hasilKali := A[2] * A[3]
	// 	fmt.Println(hasilKali)
	// }
	for i := 0; i < n; i++ {
		counter++
		fmt.Println("counter:", counter)
		if A[i] == 0 {
			return i
		}
	}
	return -1 // jika angka 0 tidak ada
}

func main() {
	data := []int{1, 2, 3, 5, 7, 9}
	fmt.Println(data[2])
	result := linier(len(data), data)
	fmt.Println("hasil", result)
}
