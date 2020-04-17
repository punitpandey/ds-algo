package main

import "fmt"

func bubble(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	return array
}

func main() {
	fmt.Println(bubble([]int{1, 3, 5, 4, 2, 9, 7}))
}
