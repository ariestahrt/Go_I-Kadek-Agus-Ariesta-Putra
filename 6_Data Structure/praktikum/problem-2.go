package main

import (
	"fmt"
)

func munculSekali(angka string) []int {
	var count = map[int]int{}
	var res = []int{}

	for _, val := range angka {
		count[int(val-'0')]++
	}

	for key, val := range count {
		if val == 1 { res = append(res, key) } 
	}
	
	return res
}

func main() {
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}