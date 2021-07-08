package main

import "fmt"

func MoveElementToEnd(array []int, toMove int) []int {

	left, right := 0, len(array)-1
	for left < right {
		if array[right] == toMove {
			right -= 1
		} else {
			if array[left] == toMove {
				array[left], array[right] = array[right], array[left]
				left += 1
				right -= 1
			} else {
				left += 1
			}
		}
	}
	return array
}

func main() {
	array := []int{2, 1, 2, 2, 2, 3, 4, 2}
	toMove := 2
	output := MoveElementToEnd(array, toMove)
	fmt.Println(output)
}
