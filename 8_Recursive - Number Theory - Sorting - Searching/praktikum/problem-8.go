package main

import "fmt"

type pair struct {
	name string
	count int
}

// Mergesort thanks to https://austingwalters.com/merge-sort-in-go-golang/
func Merge(left, right []pair) []pair {

	size, i, j := len(left)+len(right), 0, 0
	slice := make([]pair, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i].count < right[j].count {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}

func MergeSort(slice []pair) []pair {
	if len(slice) < 2 {
		return slice
	}
	mid := (len(slice)) / 2
	return Merge(MergeSort(slice[:mid]), MergeSort(slice[mid:]))
}

func MostAppearItem(items []string) []pair {
	res := []pair{}
	count := map[string]int{}

	for _, item := range items {
		if _, exist := count[item]; exist {
			count[item]++
		}else{
			count[item] = 1
		}
	}

	for key, val := range count {
		res = append(res, pair{name: key, count: val})
	}

	res_sorted := MergeSort(res)
	return res_sorted
}

func main() {
	fmt.Println(MostAppearItem([]string{"js","js","golang","ruby","ruby","js","js"}))
	fmt.Println(MostAppearItem([]string{"A","B","B","C","A","A","A", "B", "A", "D", "D"}))
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
}