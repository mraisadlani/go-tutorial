package main

import "fmt"

func main() {
	// Merubah ke tipe data lain
	// var result interface{} = random()
	// var resultString string = result.(string)

	// fmt.Println(resultString)

	var hasil interface{} = random()

	switch value := hasil.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown type")
	}
}

func random() interface{} {
	return "Vanilla"
}
