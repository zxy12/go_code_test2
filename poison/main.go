package main

import (
	"fmt"
	"math"
)

func main() {
	var bottles []int
	for i := 1; i <= 100; i++ {
		bottles = append(bottles, i)
	}

	shift := math.Ceil(math.Log2(float64(100)))
	var mousesA [][]int = make([][]int, int(shift))
	var mousesB [][]int = make([][]int, int(shift))
	for _, v := range bottles {
		for c := 0; c < int(shift); c++ {
			if v&(1<<uint(c)) > 0 {
				mousesA[c] = append(mousesA[c], v)
			} else {
				mousesB[c] = append(mousesB[c], v)
			}
		}
	}

	for i, v := range mousesA {
		fmt.Printf("mouse[%v]: ", i)
		for _, vv := range v {
			fmt.Printf("%v,", vv)
		}
		println()
	}

	println()
	var poisonedNum = 100
	fmt.Printf("%v,%b\n", poisonedNum, poisonedNum)

	for i := len(mousesA) - 1; i >= 0; i-- {
		for _, v := range mousesA[i] {
			if v == poisonedNum {
				fmt.Printf("mouseA[%v]", (i + 1))
				break
			}
		}
	}
	fmt.Println("is dead")

	for i := len(mousesB) - 1; i >= 0; i-- {
		for _, v := range mousesB[i] {
			if v == poisonedNum {
				fmt.Printf("mouseB[%v]", (i + 1))
				break
			}
		}
	}
	fmt.Println("is dead")

}
