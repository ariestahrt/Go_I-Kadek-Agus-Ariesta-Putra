package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	if b > a { return Compare(b,a) }

	if strings.Contains(a, b) == false { return Compare(a, b[1:]) }
	return b
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))     // AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  // KANG
	fmt.Println(Compare("KI", "KIJANG"))      // KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    // ILA
	fmt.Println(Compare("ILALANG", "LA"))    // LA
}