package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Vanilla",
		"address": "Jakarta",
	}

	person["job"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["job"])

	book := make(map[string]string)
	book["title"] = "E-book Go Language"
	book["author"] = "Muhammad Rais Adlani"
	book["wrong"] = "Ups"

	delete(book, "wrong")

	fmt.Println(book)
	fmt.Println(len(book))
}
