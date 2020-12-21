package main

import "fmt"

func endApp() {
	message := recover()
	fmt.Println("Terjadi error", message)

	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("APLIKASI ERROR")
	}

	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
	fmt.Println("Rais")
}
