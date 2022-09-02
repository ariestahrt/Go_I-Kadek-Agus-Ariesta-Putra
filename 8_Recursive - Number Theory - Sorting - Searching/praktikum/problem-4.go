package main

import (
	"fmt"
	"math"
)

func MaxSequence(arr []int) int {
	if len(arr) < 1 { return 0 }
	
	sum := make([]float64, len(arr) +1)
	max_sum := 0.0

	for idx, val := range arr {
		sum[idx+1] = float64(val) + sum[idx]
		max_sum = math.Max(max_sum, sum[idx+1])
	}

	// Recursive
	return int(math.Max(max_sum, float64(MaxSequence(arr[1:]))))
}

func main() {
	fmt.Println(MaxSequence([]int{-2,1,-3,4,-1,2,1,-5,4}))
	fmt.Println(MaxSequence([]int{-2,-5,6,-2,-3,1,5,-6}))
	fmt.Println(MaxSequence([]int{-2,-3,4,-1,-2,1,5,-3}))
	fmt.Println(MaxSequence([]int{-2,-5,6,-2,-3,1,6,-6}))
	fmt.Println(MaxSequence([]int{-2,-5,6,2,-3,1,6,-6}))
}