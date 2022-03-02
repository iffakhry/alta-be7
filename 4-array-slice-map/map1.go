package main

import (
	"fmt"
)

func main() {

	// var monthsArray =[12]int{30,28,31,30}

	//membuat map dengan value map
	var data = map[string]map[string]int{}
	// var temp = map[string]int{"tempA": 30}
	data["a"] = map[string]int{"tempA": 50, "tempB": 100}
	fmt.Println(data["a"]["tempB"])

	//contoh map
	var months = map[string]int{}
	months["januari"] = 30
	months["februari"] = 28
	months["maret"] = 31
	months["januari"] = 31

	fmt.Println(months["januari"])
	fmt.Println(len(months))

	months["april"] = 30
	fmt.Println(len(months))

	//akses tiap value di map
	for key, value := range months {
		fmt.Println(key, " --> ", value)
	}

	//contoh map dengan key int dan value boolean
	primes := map[int]bool{}
	primes[2] = true
	primes[3] = true
	primes[6] = false
	fmt.Println(primes[4])
	value, exist := primes[4] // check key is exist
	fmt.Println(value, exist)
	//exist = false
	// if true {}
	if exist == true {
		fmt.Println("key terdaftar")
	} else {
		fmt.Println("key belum terdaftar")
	}
}
