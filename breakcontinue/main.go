package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Perulangan ke", i)

		if i == 5 {
			break
		}
	}

	for j := 0; j < 10; j++ {
		if j%2 == 0 {
			continue
		}

		fmt.Println("Perulangan ke ", j)
	}
}
