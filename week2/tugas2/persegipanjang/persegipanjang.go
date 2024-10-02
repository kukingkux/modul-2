package main

import "fmt"

func main() {
	var p, l, K, L int
	fmt.Println("Masukkan p dan l:")
	fmt.Scan(&p, &l)
	L = p * l
	K = p*2 + l*2
	fmt.Print(K, L)
}
