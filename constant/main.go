package main

import "fmt"

func main() {
	const (
		firstname string = "Muhammad Rais"
		lastname         = "Adlani"
		value            = 1000
	)

	fmt.Println(firstname)
	fmt.Println(lastname)
	fmt.Println(value)

	// error
	// firstname = "Tidak bisa dirubah"
	// lastname = "Tidak bisa dirubah"
}
