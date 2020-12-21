package main

import "fmt"

func main() {
	// Penjumlahan
	var a = 10
	var b = 5
	var c = a + b

	fmt.Println(c)

	// Perkalian
	var a1 = 10
	var b1 = 5
	var c1 = a1 * b1

	fmt.Println(c1)

	// Pengurangan
	var a2 = 10
	var b2 = 5
	var c2 = a2 - b2

	fmt.Println(c2)

	// Pembagian
	var a3 = 10
	var b3 = 5
	var c3 = a3 / b3

	fmt.Println(c3)

	// Sisa Bagi
	var a4 = 10
	var b4 = 10
	var c4 = a4 % b4

	fmt.Println(c4)

	// Augmented Assignment
	var i = 10
	i += 10 // i = i + 10
	fmt.Println(i)

	var i1 = 10
	i1 -= 10 // i = i - 10
	fmt.Println(i1)

	var i2 = 10
	i2 *= 10 // i = i * 10
	fmt.Println(i2)

	var i3 = 10
	i3 /= 10 // i = i / 10
	fmt.Println(i3)

	var i4 = 10
	i4 %= 10 // i = i % 10
	fmt.Println(i4)

	// Unary
	var min = -100
	var plus = +100

	fmt.Println(min)
	fmt.Println(plus)

	var check = true

	fmt.Println(!check) // false

	var y = 10
	y++
	fmt.Println(y)

	var z = 10
	z--
	fmt.Println(z)
}
