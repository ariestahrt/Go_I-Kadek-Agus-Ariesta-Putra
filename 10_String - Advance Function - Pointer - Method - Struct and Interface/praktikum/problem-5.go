package main

import (
	"fmt"
	"math"
)

type Student struct {
	name  []string
	score []int
}

func (s Student) Avarage() float64 {
	sum := 0
	for _, val := range s.score {
		sum += val
	}
	return float64(float64(sum) / float64(len(s.score)))
}

func (s Student) Min() (min int, name string) {
	min = math.MaxInt32  
	for idx, val := range s.score {
		if val < min { min = val ; name = s.name[idx] }
	}
	return
}

func (s Student) Max() (max int, name string) {
	max = math.MinInt32  
	for idx, val := range s.score {
		if val > max { max = val ; name = s.name[idx] }
	}
	return
}

func main() {
	var a = Student{}
	for i := 1; i < 6; i++ {
		var name string
		fmt.Print("Input " + string(i+'0') + " Student’s Name : ")
		fmt.Scan(&name)
		a.name = append(a.name, name)
		var score int
		fmt.Print("Input " + name + " Score : ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}
	fmt.Println("\n\nAverage Score Students is", a.Avarage())
	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Students is : "+nameMax+" (", scoreMax, ")")
	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Students is : "+nameMin+" (", scoreMin, ")")
}