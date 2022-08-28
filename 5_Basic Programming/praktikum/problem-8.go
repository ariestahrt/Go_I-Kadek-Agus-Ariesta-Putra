package main

import (
	"fmt"
)

func cetakTablePerkalian(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Printf("%d\t", i * j)
		}
		fmt.Println()
	}
}

func main() {
	var n int = 9
	fmt.Print("[!] Masukkan Nilai : ")
	fmt.Scanf("%d\n", &n)
	cetakTablePerkalian(n)
}
