package main

import "fmt"

func caesar(offset int, input string) string {
	str := []rune(input)
	for i, _ := range str {
		str[i] = 'a' + (str[i] % 'a' + rune(offset % 26)) % 26
	}
	return string(str)
}

func main() {
   fmt.Println(caesar(3, "abc"))                           // def
   fmt.Println(caesar(2, "alta"))                          // cnvc
   fmt.Println(caesar(10, "alterraacademy"))               // kvdobbkkmknowi
   fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz")) // bcdefghijklmnopqrstuvwxyza
   fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
}