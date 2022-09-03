package main

import(
	"fmt"
	"math"
)

func Max(a, b int) int {
	if a > b { return a } ; return b
}

func Min(a, b int) int {
	if a < b { return a } ; return b
}

func getMinMax(numbers ...*int) (min int, max int) {
	min = math.MaxInt32
	max = math.MinInt32

	for _, val := range numbers {
		min = Min(min, *val)
		max = Max(max, *val)
	}
	return
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)
	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("Nilai min ", min)
	fmt.Println("Nilai max ", max)
}