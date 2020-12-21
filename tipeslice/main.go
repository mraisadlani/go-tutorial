package main

import "fmt"

func main() {
	names := [...]string{
		"Vanilla",
		"Muhammad Rais",
		"Glory",
		"Chunky",
		"Els",
		"Porty",
	}

	slice := names[4:6]

	fmt.Println(slice[0])
	fmt.Println(slice[1])

	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// months[5] = "Diubah"
	// fmt.Println(slice1)

	// slice1[0] = "Mei Lagi" // Index ke 0 dari slice start 4
	// fmt.Println(months)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daySlice1 := days[5:]
	daySlice1[0] = "New Sabtu"
	daySlice1[1] = "New Minggu"

	fmt.Println(days)

	daySlice2 := append(daySlice1, "Tanggal Merah")
	daySlice2[0] = "Ups"
	fmt.Println(daySlice2)
	fmt.Println(days)

	// Make
	newSlice := make([]string, 3, 5)
	newSlice[0] = "Vanilla"
	newSlice[1] = "Coklat"
	newSlice[2] = "Strawberry"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	// Array vs Slice
	intArray := [...]int{1, 2, 3, 4, 5}
	intSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(intArray)
	fmt.Println(intSlice)
}
