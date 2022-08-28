package main

import (
	"fmt"
)

func playWithAsterik(n int) {
	for i := 1; i <= n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < i; j++ {
			fmt.Print("* ")
		}
		fmt.Print("\n")
	}
}

func main() {
	var n int = 5
	fmt.Print("[!] Masukkan Nilai : ")
	fmt.Scanf("%d\n", &n)
	playWithAsterik(n)
}
