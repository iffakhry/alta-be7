package main

import (
	"fmt"
	"reflect"
)

func main() {
	var countries [3]string

	countries[0] = "India"  // Assign a value to the first element
	countries[1] = "Canada" // Assign a value to the second element
	countries[2] = "Indo"   // Assign a value to the second element

	fmt.Println(countries[0]) // Access the first element value
	fmt.Println(countries[1]) // Access the second element value
	fmt.Println(countries[2]) // Access the second element value

	odd_numbers := [5]int{1, 3, 5, 7, 9}      // Intialized with values
	var even_numbers [5]int = [5]int{1, 2, 4} // Partial assignment

	fmt.Println(odd_numbers)
	fmt.Println(even_numbers)

	primes := [...]int{2, 3, 3, 5}

	fmt.Println(reflect.ValueOf(primes).Kind())
	fmt.Println(len(primes))

	array2 := [5]string{1: "indonesia", 3: "malaysia"}
	fmt.Println(array2)

	data := [5]int{2, 3, 5}

	// technique 1
	for index := 0; index < len(data); index++ {
		fmt.Println(data[index])
	}
	// technique 2
	for index, element := range data {
		fmt.Println(index, "=>", element)
	}
	for _, value := range data {
		fmt.Println(value)
	}
	// technique 3
	index := 0
	for range data {
		fmt.Println(data[index])
		index++
	}

	ganjil := [5]int{1, 3, 5, 7, 9}
	for _, v := range ganjil {
		nilai := v + 10
		fmt.Println("bil ganjil", nilai)
	}
}
