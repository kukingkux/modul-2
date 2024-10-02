package main

import "fmt"

func main() {
	var (
		f    float64
		x, y int
	)
	fmt.Println("Masukkan x dan y: ")
	fmt.Scan(&x, &y)
	var X float64 = float64(x)
	var Y float64 = float64(y)
	f = 1/((3*X)*X+10) + 10*Y + 7
	fmt.Println(f)
}
