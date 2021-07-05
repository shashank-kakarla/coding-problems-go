package main

import "sort"

func TwoNumberSum(array []int, target int) []int {
	sort.Ints(array)
	leftIndex := 0
	rightIndex := len(array) - 1
	for leftIndex < rightIndex {
		sum := array[leftIndex] + array[rightIndex]
		if sum == target {
			return []int{array[leftIndex], array[rightIndex]}
		} else if sum < target {
			leftIndex += 1
		} else {
			rightIndex -= 1
		}
	}
	return []int{}
}
