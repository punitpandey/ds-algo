package main

import "fmt"

func merge(array []int) []int {
	if len(array) == 1 {
		return array
	}
	middle := len(array) / 2
	left := array[:middle]
	right := array[middle:]

	array = mergesort(array, merge(left), merge(right))
	return array
}

func mergesort(array []int, left []int, right []int) []int {
	var response []int
	lIndex, rIndex := 0, 0
	for lIndex < len(left) || rIndex < len(right) {
		if lIndex < len(left) && (rIndex == len(right) || left[lIndex] < right[rIndex]) {
			response = append(response, left[lIndex])
			lIndex++
		} else {
			response = append(response, right[rIndex])
			rIndex++
		}
	}
	return response
}

func main() {
	fmt.Println(merge([]int{1, 3, 5, 2, 4, 3, 9, 7}))
}
