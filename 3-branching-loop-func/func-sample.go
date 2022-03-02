package main

import "fmt"

func TambahAngka() {
	fmt.Println(1 + 1)
}

func TambahAngkaWithReturn() int {
	hasil := 10
	return hasil
}

func TambahAngkaWithParamReturn(a int) int {
	hasil := a + 10
	return hasil
}

func TambahDataWithParamReturn(a int) float64 {
	hasil := a + 10
	var data float64
	data = float64(hasil)
	return data
}

func main() {
	TambahAngka()
	result := TambahAngkaWithReturn()
	fmt.Println(result)
	data := TambahAngkaWithParamReturn(5)
	fmt.Println(data)
	fmt.Println(TambahAngkaWithParamReturn(10))
}
