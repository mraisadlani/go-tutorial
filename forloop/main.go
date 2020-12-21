package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	for count := 1; count <= 10; count++ {
		fmt.Println("Perulangan ke ", count)
	}

	slice := []string{"Vanilla", "Joko", "Rudy", "Endy"}

	for i := 0; i < len(slice); i++ {
		fmt.Println("Nama Anda ", slice[i])
	}

	names := []string{"Vanilla", "Joko", "Rudy", "Endy"}

	for _, name := range names {
		// fmt.Println("Index ke ", i, "=", name)
		fmt.Println(name)
	}

	person := make(map[string]string)
	person["name"] = "Eko"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
