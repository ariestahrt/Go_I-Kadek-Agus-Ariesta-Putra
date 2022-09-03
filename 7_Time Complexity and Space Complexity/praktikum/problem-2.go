package main

import "fmt"

// func pow(x, n int) int {
// 	res := 1
// 	for n > 0 {
// 		if n % 2 == 1 { res = res*x }
// 		n = n/2; x = x*x
// 	}
// 	return res
// }

func pow_recursive(base, n int) int {
	if n == 0 {return 1}
	if n % 2 == 1 { return base * pow_recursive(base, n-1)}
	return pow_recursive(base*base, n >> 1)
}
 
func main() {
   fmt.Println(pow_recursive(2, 3))  // 8
   fmt.Println(pow_recursive(5, 3))  // 125
   fmt.Println(pow_recursive(10, 2)) // 100
   fmt.Println(pow_recursive(2, 5))  // 32
   fmt.Println(pow_recursive(7, 3))  // 343
}
