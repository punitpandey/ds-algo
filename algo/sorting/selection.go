package main

import "fmt"

func selection(array []int) []int {
	for i := 0; i < len(array); i++ {
		min := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[min] {
				min = j
			}
		}
		if min != i {
			array[i], array[min] = array[min], array[i]
		}
	}
	return array
}

func main() {
	fmt.Println(selection([]int{1, 3, 5, 2, 9, 7}))
}
