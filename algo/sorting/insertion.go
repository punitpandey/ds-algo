package main

import "fmt"

func insertion(array []int) []int {
	for i := 1; i < len(array); i++ {
		if array[0] > array[i] {
			temp := array[i]
			copy(array[1:i+1], array[0:i])
			array[0] = temp
		} else {
			for j := 1; j < i; j++ {
				if array[j-1] <= array[i] && array[j] > array[i] {
					array[j], array[i] = array[i], array[j]
				}
			}
		}
	}
	return array
}
func main() {
	fmt.Println(insertion([]int{12, 3, 5, 2, 3, 9, 7}))
}
