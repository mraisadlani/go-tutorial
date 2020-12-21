package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("e([a-z])o")

	fmt.Println(regex.MatchString("vanilla"))
	fmt.Println(regex.MatchString("vanilli"))
	fmt.Println(regex.MatchString("vanille"))

	search := regex.FindAllString("vanilla vanilli vanille vanillu vanille ekvanillei", -1)
	fmt.Println(search)
}
