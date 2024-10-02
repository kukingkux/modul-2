package main

import "fmt"

func main() {
	var a, b, c, d, e, hasil int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&a, &b, &c, &d, &e)
	hasil = a + b + c + d + e
	fmt.Print("Hasil penjumlahan ", a, b, c, d, e, " adalah ", hasil)
}
