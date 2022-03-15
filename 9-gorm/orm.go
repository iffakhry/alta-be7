package main

type Person struct {
	ID   uint64
	Name string
	Age  uint32
	Sex  string
}

func main() {
	var person1 Person
	person1.ID = 1
	person1.Name = "Budi"
	person1.Age = 40
	person1.Sex = "male"

	var person2 Person
	person2.ID = 2
	person2.Name = "Yanto"
	person2.Age = 41
	person2.Sex = "male"

	var person = map[int]Person{}
	person[1] = Person{1, "Budi", 40, "male"}
	person[2] = Person{2, "Yanto", 41, "male"}

	var personSlice = []Person{}
	personSlice = append(personSlice, Person{1, "Budi", 40, "male"})
	personSlice = append(personSlice, Person{2, "Yanto", 41, "male"})
}
