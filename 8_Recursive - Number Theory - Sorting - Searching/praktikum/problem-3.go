package main

import "fmt"

const MAX_ARR = 1e6
var prime_list = []int{}

func primaSegiEmpat(high, wide, start int) {
	// Search idx
	idx_prime := 0
	for idx, val := range prime_list {
		if val > start {
			idx_prime = idx
			break
		}
	}

	sum := 0

	// Printnya kebalik di soal widenya jadi high, highnya jadi wide
	for ; wide > 0; wide-- {
		for i := 0; i<high; i++ {
			fmt.Printf("%d\t", prime_list[idx_prime])
			sum += prime_list[idx_prime]
			idx_prime++
		}
		fmt.Println()
	}
	fmt.Println(sum)
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

	primaSegiEmpat(2,3,13)
	primaSegiEmpat(5,2,1)
}
