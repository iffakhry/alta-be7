package main

import "fmt"

func main() {
	// var today int = 2
	// switch today {
	// case 1:
	// 	fmt.Printf("Today is Monday")
	// case 3:
	// 	fmt.Printf("Today is Wednesday")
	// case 2:
	// 	fmt.Printf("Today is Tuesday")
	// default:
	// 	fmt.Printf("Invalid Date")
	// }

	// switch {
	// case today == 1:
	// 	fmt.Printf("Today is Monday")
	// case today == 2 && today < 5:
	// 	fmt.Printf("Today is Tuesday")
	// default:
	// 	fmt.Printf("Invalid Date")
	// }

	value := 1
	switch value {
	case 100:
		fmt.Println(100)
		fallthrough
	case 42:
		fmt.Println(42)
	case 1:
		fmt.Println(1)
		fallthrough
	default:
		fmt.Println("default")
	}

}
