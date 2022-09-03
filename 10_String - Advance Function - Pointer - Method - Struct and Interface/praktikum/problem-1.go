package main

import (
	"fmt"
	_ "strings"
)

func Compare(a, b string) string {
   // your code here
   if a > b { return b }
   return a
}


func main() {
   fmt.Println(Compare("AKA", "AKASHI"))     // AKA
   fmt.Println(Compare("KANGOORO", "KANG"))  // KANG
   fmt.Println(Compare("KI", "KIJANG"))      // KI
   fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU
   fmt.Println(Compare("ILALANG", "ILA"))    // ILA
}