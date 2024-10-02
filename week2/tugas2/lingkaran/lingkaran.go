package main

import "fmt"

func main() {
	var r, luaslingkaran, kelilinglingkaran float64
	fmt.Print("Masukkan nilai r: ")
	fmt.Scan(&r)
	luaslingkaran = 3.14 * (r * r)
	kelilinglingkaran = 2 * 3.14 * r
	fmt.Println(luaslingkaran)
	fmt.Println(kelilinglingkaran)
}
