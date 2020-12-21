package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	rais := Man{"Rais"}
	rais.Married()

	fmt.Println(rais)
}
