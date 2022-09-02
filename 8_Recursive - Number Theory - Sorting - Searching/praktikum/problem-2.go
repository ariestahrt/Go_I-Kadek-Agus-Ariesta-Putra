package main

import "fmt"


func fibonacci(number int) int {
	if number == 0 { return 0 }
	if number == 1 { return 1 }
	return fibonacci(number-1) + fibonacci(number-2)	
}

func main() {
	fmt.Println(fibonacci(0))
	fmt.Println(fibonacci(2))
	fmt.Println(fibonacci(9))
	fmt.Println(fibonacci(10))
	fmt.Println(fibonacci(12))
}
