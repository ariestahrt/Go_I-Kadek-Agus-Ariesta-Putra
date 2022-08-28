package main

import (
	"fmt"
	"math"
)

func main(){
	var T float32
	var r float32

	fmt.Print("[!] Masukkan T : ")
	fmt.Scanf("%f\n", &T)
	fmt.Print("[!] Masukkan r : ")
	fmt.Scanf("%f\n", &r)
	
	// Kode
	Lp := 2 * math.Pi * r * (r + T)
	fmt.Printf("[!] Luas : %.2f\n", Lp)
}