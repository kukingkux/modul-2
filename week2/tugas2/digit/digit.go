package main

import "fmt"

func main() {
	var (
		x int
		// d1, d2, d3 int
	)
	fmt.Print("Masukkan digit: ")
	fmt.Scan(&x)

	d1 := x / 100
	d2 := (x % 100) / 10
	d3 := x % 10

	fmt.Printf("%d %d %d", d1, d2, d3)

	// kode dibawah dapat dijalankan namun mendapatkan exception berupa
	// Operation did not complete successfully because the file contains a virus or potentially unwanted software.
	// d1 = x / 100
	// d2 = (x % 100) / 10
	// d3 = x % 10

	// fmt.Print(d1, d2, d3)
}
