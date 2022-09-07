package main

import "fmt"

func playingDomino(cards [][]int, deck []int) interface{} {
	max:=-9999
	res:= []int{}
	deck_map := map[int]bool{deck[0]: true, deck[1]:true}

	for _, card := range cards {
		_, exist1 := deck_map[card[0]]
		_, exist2 := deck_map[card[1]]
		if exist1 || exist2 {
			if card[0] + card[1]>max {
				max = card[0] + card[1]
				res = card
			}
		}
	}

	if max >= 0 {
		return fmt.Sprintf("[%d %d]", res[0], res[1])
	}else{
		return 1
	}
}

func main() {
	fmt.Println(playingDomino([][]int{[]int{6,5}, []int{3,4}, []int{2,1}, []int{3,3}}, []int{4,3}))
	fmt.Println(playingDomino([][]int{[]int{6,5}, []int{3,3}, []int{3,4}, []int{2,1}}, []int{3,6}))
	fmt.Println(playingDomino([][]int{[]int{6,6}, []int{2,4}, []int{3,6}}, []int{5,1}))
}