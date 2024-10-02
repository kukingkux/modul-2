package main

import "fmt"

func main() {
	var mk string = "Algoritma dan Pemrgoraman"
	var kode string
	var sks int

	fmt.Print("Tuliskan Kode MK dan SKS : ")
	fmt.Scan(&kode, &sks)
	fmt.Print("Kredit MK ", kode, " - ", mk, " 1 adalah ", sks, "SKS")

	// fmt.Println("Mata Kuliah:", mk)
	// fmt.Println("Kode Matkul:", kode)
	// fmt.Println("Jumlah SKS:", sks)
	// fmt.Scan()
}
