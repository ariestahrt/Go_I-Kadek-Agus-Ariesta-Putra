package main

import (
	"fmt"
	"math"
)

func FindMinAndMax(arr []int) string {
	min := math.MaxInt32
	max := math.MinInt32
	idx_min := -1
	idx_max := -1

	for idx, val := range arr {
		if val < min { min = val; idx_min = idx }
		if val > max { max = val; idx_max = idx }
	}

	return fmt.Sprintf("min: %d index: %d max: %d index: %d", min, idx_min, max, idx_max)

}

func main() {
	fmt.Println(FindMinAndMax([]int{5,7,4,-2,-1,8}))
	fmt.Println(FindMinAndMax([]int{2,-5,-4,22,7,7}))
	fmt.Println(FindMinAndMax([]int{4,3,9,4,-21,7}))
	fmt.Println(FindMinAndMax([]int{-1,5,6,4,4,2,18}))
	fmt.Println(FindMinAndMax([]int{-2,5,-7,4,7,-20}))
}