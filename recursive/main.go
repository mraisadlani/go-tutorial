package main

import "fmt"

func main() {
	fmt.Println(factorialLoop(10))
}

func factorialLoop(value int) int {
	result := 1

	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}
