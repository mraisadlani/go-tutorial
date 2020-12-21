package main

import "fmt"

func main() {
	// Manual Array
	var names [3]string

	names[0] = "Muhammad"
	names[1] = "Rais"
	names[2] = "Adlani"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// Langsung
	var values = [3]int{
		90,
		80,
		70,
	}

	fmt.Println(values)

	var score = [4]int{
		80, 87, 90, 93,
	}

	fmt.Println(len(score))
	fmt.Println(len(values))
}
