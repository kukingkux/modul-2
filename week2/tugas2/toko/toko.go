package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Print("Masukkan modal barang x, y, z: ")
	fmt.Scan(&x, &y, &z)
	var X float64 = float64(x)
	var Y float64 = float64(y)
	var Z float64 = float64(z)
	X = X + X*5/100
	Y = Y + Y*5/100
	Z = Z + Z*5/100
	fmt.Print(X, Y, Z)

	// var x, y, z float64
	// fmt.Print("Masukkan modal barang x, y, z: ")
	// fmt.Scan(&x, &y, &z)
	// x = x * 1.05
	// y = y * 1.05
	// z = z * 1.05
	// fmt.Print(x, y, z)
}
