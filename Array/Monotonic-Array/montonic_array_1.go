package main

import "fmt"

func IsMonotonic(array []int) bool {
	incOrder, decOrder := true, true
	for i := 1; i < len(array); i++ {
		incOrder = incOrder && (array[i-1] >= array[i])
		decOrder = decOrder && (array[i-1] <= array[i])
	}
	return incOrder || decOrder
}

func main() {
	output := IsMonotonic([]int{-1, -5, -10, -1100, -1100, -1101, -1102, -9001})
	fmt.Println(output)
}
