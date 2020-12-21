package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Animal struct {
	Name string
	Umur int
}

type HashName interface {
	GetName() string
	GetAge() int
}

func (person Person) GetName() string {
	return person.Name
}

func (person Person) GetAge() int {
	return person.Age
}

func (animal Animal) GetName() string {
	return animal.Name
}

func (animal Animal) GetAge() int {
	return animal.Umur
}

func main() {
	var rais Person
	rais.Name = "Rais"
	rais.Age = 24

	SayHello(rais)

	cat := Animal{
		Name: "Kucing",
		Umur: 1,
	}

	SayHello(cat)
}

func SayHello(hashName HashName) {
	fmt.Println("Hello", hashName.GetName(), " and my age ", hashName.GetAge())
}
