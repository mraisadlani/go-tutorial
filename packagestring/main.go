package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Muhammad Rais", "Rais"))
	fmt.Println(strings.Split("Muhammad Rais", " "))
	fmt.Println(strings.ToLower("Muhammad Rais"))
	fmt.Println(strings.ToUpper("Muhammad Rais"))
	fmt.Println(strings.Trim("           Muhammad Rais                   ", " "))
	fmt.Println(strings.ReplaceAll("Rais Rais Rais", "Rais", "Vanilla"))
}
