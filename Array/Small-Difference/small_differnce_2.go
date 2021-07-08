package main

import (
	"fmt"
	"math"
	"sort"
)

func SmallestDifference(array1, array2 []int) []int {
	sort.Ints(array1)
	sort.Ints(array2)
	oneIndex, twoIndex := 0, 0
	minValue, diff := math.MaxInt32, math.MaxInt32
	diffArray := []int{}
	for oneIndex < len(array1) && twoIndex < len(array2) {
		first, second := array1[oneIndex], array2[twoIndex]
		if first < second {
			diff = array2[twoIndex] - array1[oneIndex]
			oneIndex += 1
		} else if first > second {
			diff = array1[oneIndex] - array2[twoIndex]
			twoIndex += 1
		} else {
			return []int{first, second}
		}
		if minValue > diff {
			minValue = diff
			diffArray = []int{first, second}
		}
	}

	return diffArray
}

func main() {
	output := SmallestDifference([]int{-1, 5, 10, 20, 28, 3}, []int{26, 134, 135, 15, 17})
	fmt.Println(output)
}
