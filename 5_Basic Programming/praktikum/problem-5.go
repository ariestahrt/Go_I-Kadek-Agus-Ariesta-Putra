package main

import (
	"fmt"
)

func palindrome(input string) bool {
	for i := 0; i<len(input) / 2; i++ {
		if input[i] != input[len(input)-1-i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(palindrome("civic"))
	fmt.Println(palindrome("katak"))
	fmt.Println(palindrome("kasur rusak"))
	fmt.Println(palindrome("mistar"))
	fmt.Println(palindrome("lion"))
}
