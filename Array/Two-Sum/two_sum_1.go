package main

func TwoNumberSum(array []int, target int) []int {
	array_map := map[int]int{}
	for _, num := range array {
		diff := target - num
		if _, found := array_map[diff]; found {
			return []int{num, diff}
		}
		array_map[num] = diff
	}

	return []int{}
}
