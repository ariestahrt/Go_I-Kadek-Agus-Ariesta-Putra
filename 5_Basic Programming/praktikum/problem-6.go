package main

import (
	"fmt"
)

func Pangkat(base int, pangkat int) int {
	if pangkat == 0 { return 1 }
	return base * Pangkat(base, pangkat-1)
}

func main() {
	fmt.Println(Pangkat(2,3))
	fmt.Println(Pangkat(5,3))
	fmt.Println(Pangkat(10,2))
	fmt.Println(Pangkat(2,5))
	fmt.Println(Pangkat(7,3))
}