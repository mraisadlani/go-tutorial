package main

import "fmt"

func main() {
	nilai := 80

	switch {
	case nilai > 90:
		fmt.Println("A")
	case nilai > 85 && nilai < 90:
		fmt.Println("B")
	case nilai > 70 && nilai < 85:
		fmt.Println("C")
	case nilai > 60 && nilai < 70:
		fmt.Println("D")
	case nilai > 50 && nilai < 60:
		fmt.Println("E")
	case nilai < 50:
		fmt.Println("F")
	default:
		fmt.Println("Ada kesalahan input")
	}

	name := "Vanilla"
	switch name {
	case "Vanilla":
		fmt.Println("Hello Vanilla")
	case "Rudy":
		fmt.Println("Hello Rudy")
	case "Joko":
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Anda siapa ?")
	}
}
