package main

import "fmt"

func ubahValue(data string) string {
	// data = "mike"
	var point *string = &data
	*point = "mike"
	return data
}

//pass by reference
func ubahValuePointer(data *string) string {
	*data = "doe"
	return *data
}

func main() {
	var kata = "john"
	var kataResult = ubahValue(kata)
	fmt.Println(kataResult) //mike
	fmt.Println(kata)       //john

	fmt.Println("=================================")
	var kataResultPointer = ubahValuePointer(&kata)
	fmt.Println(kataResultPointer) //doe
	fmt.Println(kata)              //doe
}
