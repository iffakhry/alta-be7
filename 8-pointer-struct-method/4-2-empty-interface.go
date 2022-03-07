package main

import (
	"fmt"
	"strings"
)

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)

	i = true
	describe(i)

	i = []int{1, 2, 3}
	describe(i)

	fmt.Println([]int{1, 2, 3})

	var data map[string]interface{} = map[string]interface{}{}
	data["satu"] = "pertama"
	data["dua"] = 2
	fmt.Println(data)

	//type assertion
	var secret interface{}

	secret = 2
	var number = secret.(int) * 10
	fmt.Println(secret, "multiplied by 10 is :", number)

	secret = []string{"apple", "manggo", "banana"}
	var gruits = strings.Join(secret.([]string), ", ")
	fmt.Println(gruits, "is my favorite fruits")

	//type switch
	explain("Hello World")
	explain(52)
	explain(true)

}

func explain(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("i stored string ", strings.ToUpper(i.(string)))
	case int:
		fmt.Println("i stored int", i)
	default:
		fmt.Println("i stored something else", i)
	}
}
