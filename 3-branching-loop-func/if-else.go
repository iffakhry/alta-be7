package main

import "fmt"

func main() {
	myAge := 10
	// var age = 17
	if myAge < 5 {
		fmt.Println("too young")
	} else if myAge >= 5 && myAge <= 15 {
		fmt.Println("myage between 5 and 15")
	} else if myAge == 10 {
		fmt.Println("myage 10")
	} else {
		fmt.Println("old")
	}
	if myAge < 5 {
		fmt.Println("too young")
	}
	if myAge >= 5 && myAge <= 15 {
		fmt.Println("myage between 5 and 15")
	}
	if myAge == 10 {
		fmt.Println("myage 10")
	}

	// if myAge >= 5 && myAge <= 15 {
	// 	fmt.Println("myage between 5 and 15")
	// }

	// if dadAge := 5; dadAge > myAge {
	// 	fmt.Println("my dad")
	// }

}
