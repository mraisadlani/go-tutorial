package main

import "fmt"

func main() {
	runApplication(10)
}

func runApplication(value int) {
	defer Logging()
	fmt.Println("Run Application")

	result := 10 / value
	fmt.Println("Result", result)
}

func Logging() {
	fmt.Println("Selesai memanggil function")
}
