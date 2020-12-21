package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello() {
	fmt.Println("Hello, My name is", customer.Name)
}

func main() {
	// Create data struct
	var cs Customer
	cs.Name = "Muhammad Rais Adlani"
	cs.Address = "Indonesia"
	cs.Age = 30

	fmt.Println(cs)
	fmt.Println(cs.Name)
	fmt.Println(cs.Address)
	fmt.Println(cs.Age)

	// Create struct literals
	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     30,
	}

	fmt.Println(joko)

	budi := Customer{"Budi", "Indonesia", 30}
	fmt.Println(budi)

	rully := Customer{Name: "Rully"}
	rully.sayHello()
}
