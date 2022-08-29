package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	var flag = map[string]bool{}
	var res = []string{}

	for _, val := range append(arrayA, arrayB...) {
		if _, exist := flag[val]; exist { continue }
		res = append(res, val)
		flag[val] = true
	}
	return res
}

func main() {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))

	fmt.Println(ArrayMerge([]string{}, []string{}))
}