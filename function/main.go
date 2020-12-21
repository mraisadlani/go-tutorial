package main

import "fmt"

type Filter func(string) string

func main() {
	sayHello("Muhammad Rais", "Adlani")

	getHello := getHello("Vanilla")
	fmt.Println(getHello)

	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	first, middle, last := getNickname()
	fmt.Println(first, middle, last)

	// Function value
	goodbye := getGoodBye
	fmt.Println(goodbye("Vanilla"))

	sayHelloWithFilter("Rais", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)
}

func sayHello(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func getHello(name string) string {
	return "Hello " + name
}

func getFullName() (string, string) {
	return "Vanilla", "JS"
}

func getNickname() (firstName, middleName, lastName string) {
	firstName = "Muhammad"
	middleName = "Rais"
	lastName = "Adlani"

	return
}

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func sayHelloWithFilter(name string, filter Filter) {
	nameFilter := filter(name)
	fmt.Println("Hello", nameFilter)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}
