package main

import (
	"fmt"
	"strings"
)

const (
	str    = "something"
	substr = "somet"
)

func main() {
	// 1. len string
	sentence := "Hello"
	lenSentence := len(sentence)
	fmt.Println(lenSentence)

	// 2. compare string
	str1 := "abc"
	str2 := "abc"
	fmt.Println(str1 == str2)

	// 3. Contains
	res := strings.Contains(str, substr)
	fmt.Println(res) // true

	// 4. substring
	value := "cat;dog"
	// Take substring from index 4 to length of string.
	substring := value[:4]
	fmt.Println(substring)
	// 5. Replace
	s := "this[th[in[gs]I would like to remove"
	t := strings.Replace(s, "[", "-", -1)
	fmt.Printf("%s\n", t)

	// 6. Insert
	p := "green"
	index := 2
	q := p[:index] + "HI" + p[index:]
	fmt.Println(p, q)
}
