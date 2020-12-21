package main

import "fmt"

func main() {
	name := "Vanilla"

	if name == "Vanilla" {
		fmt.Println("Hello Vanilla")
	} else if name == "Joko" {
		fmt.Println("Hello Joko")
	} else if name == "Budi" {
		fmt.Println("Hello Budi")
	} else {
		fmt.Println("Hi, boleh kenalan ?")
	}

	var length = len(name)
	if length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

	// atau

	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
