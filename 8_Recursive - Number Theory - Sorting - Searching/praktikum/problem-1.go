package main

import "fmt"

const MAX_ARR = 1e6
var prime_list = []int{}

func primeX(number int) int {
	return prime_list[number-1]
}

func main() {
	// GENERATE PRIMES USING SOE
	not_prime := [MAX_ARR]bool {true, true}
	for i:=2; i*i < MAX_ARR; i++ {
		if not_prime[i] == false {
			for j:=i*i; j<MAX_ARR; j+=i {
				not_prime[j] = true
			}
		}
	}

	for idx, val := range not_prime {
		if val == false { prime_list = append(prime_list, idx) }
	}

	// END OF SOE

	fmt.Println(primeX(1))
	fmt.Println(primeX(5))
	fmt.Println(primeX(8))
	fmt.Println(primeX(9))
	fmt.Println(primeX(10))
}
