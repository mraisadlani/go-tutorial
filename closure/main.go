package main

import "fmt"

func main() {
	name := "Vanilla"
	counter := 0

	increment := func() {
		name = "Rais"
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
