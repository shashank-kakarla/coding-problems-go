package main

import (
	"fmt"
	"math"
)

func SmallestDifference(array1, array2 []int) []int {
	array3 := []int{}
	min := math.MaxFloat64
	for _, i := range array1 {
		for _, j := range array2 {
			if sum := math.Abs(float64(i) - float64(j)); min > sum {
				min = sum
				array3 = []int{i, j}
			}
		}
	}
	return array3
}

func main() {
	output := SmallestDifference([]int{-1, 5, 10, 20, 28, 3}, []int{26, 134, 135, 15, 17})
	fmt.Println(output)
}
