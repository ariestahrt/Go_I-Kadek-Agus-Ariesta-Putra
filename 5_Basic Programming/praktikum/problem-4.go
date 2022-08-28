package main

import (
	"fmt"
	"math"
)

func primeNumber(number int) bool {
	if number < 2 { return false }

	for i := 2; i<= int(math.Sqrt(float64(number))); i++ {
		if number % i == 0 {
			return false
		} 
	}
	return true
}

func main() {
	fmt.Println(primeNumber(11))
	fmt.Println(primeNumber(13))
	fmt.Println(primeNumber(17))
	fmt.Println(primeNumber(20))
	fmt.Println(primeNumber(35))
}
