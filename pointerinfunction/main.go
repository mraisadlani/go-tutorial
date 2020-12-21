package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var alamat = Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "",
	}

	ChangeCountryToIndonesia(&alamat)

	fmt.Println(alamat)
}

func ChangeCountryToIndonesia(alamat *Address) {
	alamat.Country = "Indonesia"
}
