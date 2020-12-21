package main

import "fmt"

func main() {
	total := sumAll(10, 10, 10, 10, 10, 10)
	fmt.Println(total)

	// Using slice
	numbers := []int{10, 10, 10, 10, 10}
	hasil := sumAll(numbers...)
	fmt.Println(hasil)
}

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}
