package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	if b > a { return Compare(b,a) }	
	if strings.Contains(a,b){ return b }
	
	cutLeft  := Compare(a,b[1:])
	cutRight := Compare(a,b[:len(b)-1])
	
	if len(cutLeft) > len(cutRight) { return cutLeft }
	return cutRight
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))     // AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  // KANG
	fmt.Println(Compare("KI", "KIJANG"))      // KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    // ILA
	fmt.Println(Compare("ILALANG", "LA"))    // LA
	fmt.Println(Compare("MAKANNASI", "NANNAB"))    // ANNA
}