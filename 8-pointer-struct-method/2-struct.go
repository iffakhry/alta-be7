package main

import (
	"fmt"
)

type Person struct {
	Firstname string
	Lastname  string
	Age       uint
}

func main() {
	// var person1firstname string = "john"
	// var person1lastname string = "doe"

	var person1 Person
	person1.Firstname = "john"
	person1.Lastname = "doe"
	person1.Age = 20

	fmt.Println(person1.Lastname) //doe

	var person2 Person
	person2.Firstname = "mike"
	person2.Lastname = "tyson"
	person2.Age = 30

	fmt.Println(person2.Lastname) //tyson

	type Student struct {
		Name    string
		NIK     string
		Age     uint
		Address string
		BB      float32
	}

	var student1 Student
	student1.Name = "mike"
	student1.NIK = "0123456"
	student1.Address = "jakarta"
	student1.BB = 60.8
	fmt.Println(student1)

	var student2 = Student{Name: "john", NIK: "09762", Age: 17, Address: "malang", BB: 50.1}
	fmt.Println(student2)

	var peserta = map[string]Student{}
	peserta["john"] = Student{Name: "john", NIK: "09762", Age: 17, Address: "malang", BB: 50.1}
	peserta["doe"] = Student{Name: "doe", NIK: "1111", Age: 18, Address: "jakarta", BB: 52.1}
	fmt.Println(peserta["doe"].NIK)

	var pesertaslice []Student
	pesertaslice = append(pesertaslice, Student{Name: "john", NIK: "09762", Age: 17, Address: "malang", BB: 50.1})
	pesertaslice = append(pesertaslice, Student{Name: "john1", NIK: "19762", Age: 16, Address: "jakarta", BB: 50.1})
	fmt.Println(pesertaslice)
}
