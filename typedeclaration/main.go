package main

import "fmt"

func main() {
	type noKTP string
	type Married bool

	var ktpRais noKTP = "1232152675647657654"
	fmt.Println(ktpRais)

	var marriedStatus Married = true
	fmt.Println(marriedStatus)
}
