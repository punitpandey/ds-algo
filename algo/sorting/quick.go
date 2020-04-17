package main

import (
	"fmt"
	"math/rand"
)

func quick(array []int) (response []int) {
	if len(array) <= 1 {
		return array
	}
	left, right := 0, len(array)-1
	pivot := rand.Intn(len(array))
	array[right], array[pivot] = array[pivot], array[right]
	for i := 0; i < len(array); i++ {
		if array[i] < array[right] {
			array[left], array[i] = array[i], array[left]
			left++
		}
	}
	array[left], array[right] = array[right], array[left]
	quick(array[:left])
	quick(array[left+1:])
	return array
}

func main() {
	fmt.Println(quick([]int{1, 8, 5, 9, 7, 2, 4}))
}
